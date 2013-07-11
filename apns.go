package apns

import (
	"encoding/json"
)

type Custom map[string]interface{}

type Notification struct {
	Alert        string
	Badge        int
	Sound        string
	ActionLocKey string
	LaunchImage  string
	Custom
}

type LocNotification struct {
	Args         []string
	Alert        string
	Badge        int
	Sound        string
	ActionLocKey string
	LaunchImage  string
	Custom
}

func NewNotification() *Notification {
	// If someone is using the `NewNotification()` style, we should pre-make the maps for them.
	return &Notification{
		Custom: Custom{},
	}
}

func NewLocNotification() *LocNotification {
	// If someone is using the `NewLocNotification()` style, we should pre-make the maps for them.
	this := new(LocNotification)
	this.Custom = Custom{}

	return this
}

func (this *LocNotification) ToString() (string, error) {
	// I don't like going from Struct to Map to JSON, but this is the best solution
	// I can come up with right now to continue keeping the API simple and elegant.
	payload := make(map[string]interface{})
	aps := make(map[string]interface{})

	// If ActionLocKey or LaunchImage are set then we need to use the alert dictionary format.
	if this.ActionLocKey != "" || this.LaunchImage != "" || len(this.Args) > 0 {
		alert := make(map[string]interface{})

		if this.Alert != "" {
			alert["body"] = this.Alert
		}

		if this.ActionLocKey != "" {
			alert["action-loc-key"] = this.ActionLocKey
		}

		if this.LaunchImage != "" {
			alert["launch-image"] = this.LaunchImage
		}

		if len(this.Args) > 0 {
			alert["loc-args"] = this.Args
		}

		aps["alert"] = &alert
	} else if this.Alert != "" {
		aps["alert"] = this.Alert
	}

	// The omitempty option of `json.Marshal` considers "0" to be an empty value.
	// Although it's not documented, you can send Apple "-1" instead of "0" in order
	// to clear the badge icon, which saves us from needing to write our own omitempty function.
	aps["badge"] = this.Badge
	if this.Badge == 0 {
		aps["badge"] = -1
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

	bytes, err := json.MarshalIndent(payload, "", "  ")
	return string(bytes), err
}

func (this *Notification) ToString() (string, error) {
	// I don't like going from Struct to Map to JSON, but this is the best solution
	// I can come up with right now to continue keeping the API simple and elegant.
	payload := make(map[string]interface{})
	aps := make(map[string]interface{})

	// If ActionLocKey or LaunchImage are set then we need to use the alert dictionary format.
	if this.ActionLocKey != "" || this.LaunchImage != "" {
		alert := make(map[string]interface{})

		if this.Alert != "" {
			alert["body"] = this.Alert
		}

		if this.ActionLocKey != "" {
			alert["action-loc-key"] = this.ActionLocKey
		}

		if this.LaunchImage != "" {
			alert["launch-image"] = this.LaunchImage
		}

		aps["alert"] = &alert
	} else if this.Alert != "" {
		aps["alert"] = this.Alert
	}

	// The omitempty option of `json.Marshal` considers "0" to be an empty value.
	// Although it's not documented, you can send Apple "-1" instead of "0" in order
	// to clear the badge icon, which saves us from needing to write our own omitempty function.
	aps["badge"] = this.Badge
	if this.Badge == 0 {
		aps["badge"] = -1
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

	bytes, err := json.MarshalIndent(payload, "", "  ")
	return string(bytes), err
}
