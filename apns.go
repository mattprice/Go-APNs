package apns

import (
	"encoding/json"
	"fmt"
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
	case int:
		aps["badge"] = this.Badge
	default:
		err := fmt.Errorf("The badge count should be of type `int` but we found a `%T` instead.", this.Badge)
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

	// If `toPayload()` resulted in an error, send it instead of trying to generate JSON.
	if err != nil {
		return nil, err
	} else {
		return json.Marshal(payload)
	}
}

func (this *Notification) ToString() (string, error) {
	payload, err := this.toPayload()

	// If `toPayload()` resulted in an error, send it instead of trying to generate JSON.
	if err != nil {
		return "", err
	} else {
		bytes, err := json.MarshalIndent(payload, "", "  ")
		return string(bytes), err
	}
}
