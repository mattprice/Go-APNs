package apns

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

const (
	// The maximum size allowed for a notification payload is 256 bytes.
	// Any notification above this limit are refused.
	MAX_PAYLOAD_SIZE = 265
)

type Custom map[string]interface{}

type Notification struct {
	Alert        string
	AlertLocKey  string
	AlertLocArgs []string
	ActionLocKey string
	Badge        interface{}
	Sound        string
	LaunchImage  string
	Custom
}

func NewNotification() *Notification {
	// If someone is using the `NewNotification()` style, we should pre-make the maps for them.
	return &Notification{
		Custom: Custom{},
	}
}

func (this *Notification) toPayload() (*map[string]interface{}, error) {
	// I don't like going from Struct to Map to JSON, but this is the best solution
	// I can come up with right now to continue keeping the API simple and elegant.
	payload := make(map[string]interface{})
	aps := make(map[string]interface{})

	// There's 3 cases in which we might need to use the alert dictionary format.
	//	1) A localized action key is set (ActionLocKey).
	// 2) A localized alert key is set (AlertLocKey).
	//	3) A custom launch image is set (LaunchImage).
	if this.ActionLocKey != "" || this.AlertLocKey != "" || this.LaunchImage != "" {
		alert := make(map[string]interface{})

		// Don't send a body if there is a localized alert key set.
		// TODO: Log a warning about the value of `this.Alert` being ignored.
		if this.Alert != "" && this.AlertLocKey == "" {
			alert["body"] = this.Alert
		}

		if this.ActionLocKey != "" {
			alert["action-loc-key"] = this.ActionLocKey
		}

		if this.LaunchImage != "" {
			alert["launch-image"] = this.LaunchImage
		}

		if this.AlertLocKey != "" {
			alert["loc-key"] = this.AlertLocKey

			// This check is nested because you can send an alert key without
			// sending any arguments, but not the otherway around.
			if len(this.AlertLocArgs) > 0 {
				alert["loc-args"] = this.AlertLocArgs
			}
		}

		aps["alert"] = &alert
	} else if this.Alert != "" {
		aps["alert"] = this.Alert
	}

	// We use an `interface{}` for `this.Badge` because the `int` type is always initalized to 0.
	// That means we wouldn't be able to tell if someone had explicitly set `this.Badge` to 0
	// or if they had not set it at all. This switch checks let's us make sure it was
	// set explicitly, and to an integer, before storing it in the payload.
	switch this.Badge.(type) {
	case nil:
		// If we don't check for the nil case (no badge set), then default will catch it.
		break
	case int:
		aps["badge"] = this.Badge
	default:
		// TODO: Need to check and see if the badge count can be a string, too.
		err := fmt.Errorf("The badge count should be of type `int`, but we found a `%T` instead.", this.Badge)
		return nil, err
	}

	if this.Sound != "" {
		aps["sound"] = this.Sound
	}

	// All standard dictionaries need to be wrapped in the "aps" namespace.
	payload["aps"] = &aps

	// Output all the custom dictionaries.
	for key, value := range this.Custom {
		payload[key] = value
	}

	return &payload, nil
}

func (this *Notification) ToJSON() ([]byte, error) {
	payload, err := this.toPayload()
	if err != nil {
		return nil, err
	}

	return json.Marshal(payload)
}

func (this *Notification) ToString() (string, error) {
	payload, err := this.toPayload()
	if err != nil {
		return "", err
	}

	bytes, err := json.MarshalIndent(payload, "", "  ")
	return string(bytes), err
}

func (this *Notification) ToBytes() ([]byte, error) {
	// Convert the hex string iOS returns into a device token.
	// TODO: Move this into a separate `SendTo()` function.
	token, err := hex.DecodeString("19e5d3a4a27eb08e9b2d22166152a5492fd645868f1e6909e80ba99256c8590f")
	if err != nil {
		return nil, err
	}

	payload, err := this.ToJSON()
	if err != nil {
		return nil, err
	}

	// If the payload is larger than the maximum size allowed by Apple, fail with an error.
	// TODO: We should truncate the "Alert" key instead of completely bailing out. (Optional?)
	if len(payload) > MAX_PAYLOAD_SIZE {
		err := fmt.Errorf("Payload is larger than the %v byte limit.", MAX_PAYLOAD_SIZE)
		return nil, err
	}

	// Create a binary message using the new enhanced format.
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, uint8(1))             // Command
	binary.Write(buffer, binary.BigEndian, uint32(1))            // Identifier
	binary.Write(buffer, binary.BigEndian, uint32(0))            // Expiry
	binary.Write(buffer, binary.BigEndian, uint16(len(token)))   // Device token length
	binary.Write(buffer, binary.BigEndian, token)                // Token
	binary.Write(buffer, binary.BigEndian, uint16(len(payload))) // Payload length
	binary.Write(buffer, binary.BigEndian, payload)              // Payload

	return buffer.Bytes(), nil
}
