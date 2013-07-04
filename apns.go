package apns

import (
	"encoding/json"
)

type Notification struct {
	Alert        string
	Badge        int
	Sound        string
	ActionLocKey string
	LaunchImage  string
}

func (this *Notification) ToString() (string, error) {
	// I don't like going from Struct to Map to JSON, but this is the best solution
	// I can come up with right now to continue keeping the API simple and elegant.
	alert := map[string]interface{}{}

	if this.Alert != "" {
		alert["alert"] = this.Alert
	}

	// The omitempty option of `json.Marshal` considers "0" to be an empty value.
	// Although it's not documented, you can send Apple "-1" instead of "0" in order
	// to clear the badge icon, which saves us from needing to write our own omitempty function.
	alert["badge"] = this.Badge
	if this.Badge == 0 {
		alert["badge"] = -1
	}

	if this.Sound != "" {
		alert["sound"] = this.Sound
	}

	bytes, err := json.MarshalIndent(alert, "", "  ")

	return string(bytes), err
}
