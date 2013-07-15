package apns

import (
	"encoding/json"
	"testing"
)

// TODO: This is all very inefficient. It's mostly meant as a proof of concept.
func Test_Notification_ToJSON(t *testing.T) {
	// Create a basic notification.
	payload := &Notification{
		Alert: "Hello, World! This is a test.",
	}

	// Turn the notification into JSON.
	bytes, err := payload.ToJSON()
	if err != nil {
		t.Error(err)
	}

	// Turn the JSON back into a map.
	output := make(map[string]interface{})
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		t.Error(err)
	}

	// Verify the data of the map.
	aps := output["aps"].(map[string]interface{})
	if aps["alert"] != payload.Alert {
		t.Errorf("Output does not match: %#v != %#v", output["Alert"], payload.Alert)
	}
}
