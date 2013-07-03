package apns

import (
	"encoding/json"
)

type Notification struct {
	Alert string
	Badge int
}

func (this *Notification) ToString() (string, error) {
	bytes, err := json.MarshalIndent(this, "", "  ")
	return string(bytes), err
}
