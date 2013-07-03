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
	alert := &aps{this}
	bytes, err := json.MarshalIndent(alert, "", "  ")

	return string(bytes), err
}
