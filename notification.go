// Package APNs allows you to easily send Apple Push Notifications.
package apns

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

const (
	// The maximum size allowed for a notification payload is 256 bytes.
	// Any notifications larger than this limit are refused by Apple.
	MAX_PAYLOAD_SIZE = 265

	// The highest notification identifier we can send. Since the identifier can
	// only be 4 bytes long, this is the maximum value of a 32-bit unsigned integer.
	MAX_IDENTIFIER = 4294967295
)

var (
	nextIdentifier uint32 = 0
)

type Notification struct {
	// An alert message to display to the user.
	Alert string

	// The localization key for an alert message. `Alert` will be ignored if this is set.
	AlertLocKey string

	// Strings that replace the format specifiers `%@` and `%n$@` in a localized alert.
	AlertLocArgs []string

	// The localization key for the right button's title. If the value is `null`,
	// the alert will have a single "OK" button that dismisses the alert when tapped.
	// If not set, a "Close" and "View" button will be displayed.
	ActionLocKey string

	// The number to display as the badge of the application icon. To remove the
	// badge, set the value to `0`. If not set, the badge is not changed.
	Badge interface{}

	// The name of a sound file in the application bundle. If the sound file doesn't
	// exist or the value is set to "default", the default alert sound will be played.
	Sound string

	// The name of an image file in the application bundle. If not set, iOS will use
	// the previous snapshot, the image identified by the `UILaunchImageFile` key in
	// the application's `Info.plist` file, or the `Default.png`.
	LaunchImage string

	// A Unix timestamp identifying when the notification is no longer valid and
	// can be discarded by the Apple servers if not yet delivered.
	Expiry int64

	// Sets whether this should be sent via a sandbox connection or a production connection.
	// If you have not loaded the appropriate certificates, this will fail.
	Sandbox bool

	// Custom values your app can use to set context for the user interface.
	// You should not include customer information or any sensitive data.
	Custom
}

type Custom map[string]interface{}

// NewNotification creates and returns a new notification with all child maps
// and structures pre-initialized.
func NewNotification() *Notification {
	// Pre-make any required maps or other structures.
	return &Notification{
		Custom: Custom{},
	}
}

// SetExpiry accepts a Unix timestamp that identifies when the notification
// is no longer valid and can be discarded by the Apple servers if not yet delivered.
func (this *Notification) SetExpiry(expiry int64) {
	this.Expiry = expiry
}

// SetExpiryTime accepts a `time.Time` that identifies when the notification
// is no longer valid and can be discarded by the Apple servers if not yet delivered.
func (this *Notification) SetExpiryTime(t time.Time) {
	this.Expiry = t.Unix()
}

// SetExpiryDuration accepts a `time.Duration` that identifies when the notification
// is no longer valid and can be discarded by the Apple servers if not yet delivered.
// The Duration given will be added to the result of `time.Now()`.
func (this *Notification) SetExpiryDuration(d time.Duration) {
	t := time.Now().Add(d)
	this.Expiry = t.Unix()
}

// toPayload converts a Notification into a map capable of being marshaled into JSON.
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

// ToJSON generates compact JSON from a notification payload.
func (this *Notification) ToJSON() ([]byte, error) {
	payload, err := this.toPayload()
	if err != nil {
		return nil, err
	}

	return json.Marshal(payload)
}

// ToString generates indented, human readable JSON from a notification payload.
func (this *Notification) ToString() (string, error) {
	payload, err := this.toPayload()
	if err != nil {
		return "", err
	}

	bytes, err := json.MarshalIndent(payload, "", "  ")
	return string(bytes), err
}

// ToBytes converts a JSON payload into a binary format for transmitting to Apple's
// servers over a socket connection.
func (this *Notification) ToBytes(token []byte) ([]byte, error) {
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
	binary.Write(buffer, binary.BigEndian, nextIdentifier)       // Identifier
	binary.Write(buffer, binary.BigEndian, uint32(this.Expiry))  // Expiry
	binary.Write(buffer, binary.BigEndian, uint16(len(token)))   // Device token length
	binary.Write(buffer, binary.BigEndian, token)                // Token
	binary.Write(buffer, binary.BigEndian, uint16(len(payload))) // Payload length
	binary.Write(buffer, binary.BigEndian, payload)              // Payload

	// If the next identifier is greater than the max identifier, reset it.
	if nextIdentifier >= MAX_IDENTIFIER {
		nextIdentifier = 0
	}
	nextIdentifier++

	return buffer.Bytes(), nil
}

func (this *Notification) SendTo(token string) error {
	// Convert the hex string iOS returns into a device token.
	byteToken, err := hex.DecodeString(token)
	if err != nil {
		return err
	}

	payload, err := this.ToBytes(byteToken)
	if err != nil {
		return err
	}

	if this.Sandbox {
		_, err = sandboxConnection.Write(payload)
		if err != nil {
			return err
		}

		fmt.Println(sandboxConnection.ReadErrors())
	} else {
		_, err = productionConnection.Write(payload)
		if err != nil {
			return err
		}

		fmt.Println(productionConnection.ReadErrors())
	}

	return nil
}

// DebugBinary outputs each portion of the binary enhanced format for manual verification.
func (this *Notification) DebugBinary(token string) error {
	// TODO: This duplicates some of the SendTo function. Not sure what we want to do about that.

	byteToken, err := hex.DecodeString(token)
	if err != nil {
		return err
	}

	output, err := this.ToBytes(byteToken)
	if err != nil {
		return err
	}

	// Convert the Expiry to a string.
	var expiry uint32
	buffer := bytes.NewBuffer(output[5:9])
	binary.Read(buffer, binary.BigEndian, &expiry)

	// Convert the Identifier to a string.
	var identifier uint32
	buffer = bytes.NewBuffer(output[1:5])
	binary.Read(buffer, binary.BigEndian, &identifier)

	// Convert the Token Length to a string.
	var tokenLength uint32
	buffer = bytes.NewBuffer(output[9:11])
	binary.Read(buffer, binary.BigEndian, &tokenLength)

	// Convert the Token Length to a string.
	var payloadLength uint32
	buffer = bytes.NewBuffer(output[43:45])
	binary.Read(buffer, binary.BigEndian, &payloadLength)

	fmt.Println("Binary Output:")
	fmt.Println("- Command:\t", output[0])
	fmt.Println("- Identifier:\t", identifier)
	fmt.Println("- Expiry:\t", expiry)
	fmt.Println("- Token Len:\t", tokenLength)
	fmt.Println("- Token:\t", hex.EncodeToString(output[11:43]))
	fmt.Println("- Paylod Len:\t", payloadLength)
	fmt.Println("- Payload:\t", string(output[45:]))

	return nil
}
