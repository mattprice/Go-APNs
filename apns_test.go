package apns

import (
	"testing"
)

func Test_Notification_ToJSON(t *testing.T) {
	// Test a basic notification.
	payload := NewNotification()
	payload.Alert = "Hello, World! This is a test."

	bytes, err := payload.ToJSON()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if len(bytes) != 49 {
		t.Errorf("Converting basic notification to JSON failed. Expected 49 bytes but found %v.", len(bytes))
	}

	// Test a complex notification.
	payload = NewNotification()
	payload.Alert = "Hello, World! This is a test."
	payload.Badge = 42
	payload.Sound = "Test.aiff"
	payload.LaunchImage = "Default.png"
	payload.Custom["Engines"] = Custom{
		"DDG":  "https://duckduckgo.com/",
		"Bing": "https://bing.com/",
	}

	bytes, err = payload.ToJSON()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if len(bytes) != 189 {
		t.Errorf("Converting complex notification to JSON failed. Expected 189 bytes but found %v.", len(bytes))
	}

	// Test a localized notification.
	// Conveniently, this also tests setting no badge count.
	payload = NewNotification()
	payload.Alert = "This should be ignored."
	payload.AlertLocKey = "APP_GREETING_KEY"
	payload.AlertLocArgs = []string{"John", "Smith"}

	bytes, err = payload.ToJSON()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if len(bytes) != 76 {
		t.Errorf("Converting localized notification to JSON failed. Expected 76 bytes but found %v.", len(bytes))
	}

	// Test a badge count of 0.
	// This was a problem in the past when using other `json.Marshal` methods.
	payload = NewNotification()
	payload.Badge = 0

	bytes, err = payload.ToJSON()
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if len(bytes) != 19 {
		t.Errorf("Alert with badge count of 0 failed. Expected 19 bytes but found %v.", len(bytes))
	}
}
