package apns

import (
	"encoding/json"
)

type Notification struct {
	Alert string `json:"alert,omitempty"`
	Badge int    `json:"badge,omitempty"`
	Sound string `json:"sound,omitempty"`
}

type aps struct {
	*Notification `json:"aps,omitempty"`
}

func (this *Notification) ToString() (string, error) {
	// The omitempty option of `json.Marshal` considers "0" to be an empty value.
	// Although it's not documented, you can send Apple "-1" instead of "0" in order
	// to clear the badge icon, which saves us from needing to write our own omitempty function.
	if this.Badge == 0 {
		this.Badge = -1
	}

	alert := &aps{this}
	bytes, err := json.MarshalIndent(alert, "", "  ")

	return string(bytes), err
}
