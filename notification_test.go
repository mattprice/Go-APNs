package apns

import (
	"testing"
)

func Test_Notification_ToJSON(t *testing.T) {
	// Test an empty alert.
	{
		payload := NewNotification()

		bytes, err := payload.ToJSON()
		if err != nil {
			t.Error("Unexpected error:", err)
		}

		expected := `{"aps":{}}`
		if received := string(bytes); expected != received {
			t.Errorf("Expected %#q but found %#q.", expected, received)
		}
	}

	// Test basic alerts.
	{
		payload := NewNotification()
		payload.Alert = "Hello, World!"

		bytes, err := payload.ToJSON()
		if err != nil {
			t.Error("Unexpected error:", err)
		}

		expected := `{"aps":{"alert":"Hello, World!"}}`
		if received := string(bytes); expected != received {
			t.Errorf("Expected %#q but found %#q.", expected, received)
		}
	}

	// Test the alert dictionary format.
	{
		payload := NewNotification()
		payload.Alert = "Hello, World!"
		payload.Sound = "Test.aiff"
		payload.LaunchImage = "Default.png"

		bytes, err := payload.ToJSON()
		if err != nil {
			t.Error("Unexpected error:", err)
		}

		expected := `{"aps":{"alert":{"body":"Hello, World!","launch-image":"Default.png"},"sound":"Test.aiff"}}`
		if received := string(bytes); expected != received {
			t.Errorf("Expected %#q but found %#q.", expected, received)
		}
	}

	// Test a localized notification.
	// Conveniently, this also tests setting no badge count.
	{
		payload := NewNotification()
		payload.Alert = "This should be ignored."
		payload.AlertLocKey = "APP_GREETING_KEY"
		payload.AlertLocArgs = []string{"John", "Smith"}

		bytes, err := payload.ToJSON()
		if err != nil {
			t.Error("Unexpected error:", err)
		}

		expected := `{"aps":{"alert":{"loc-args":["John","Smith"],"loc-key":"APP_GREETING_KEY"}}}`
		if received := string(bytes); expected != received {
			t.Errorf("Expected %#q but found %#q.", expected, received)
		}
	}

	// Test a badge count of 0.
	// This was a problem in the past when using other `json.Marshal` methods.
	{
		payload := NewNotification()
		payload.Badge = 0

		bytes, err := payload.ToJSON()
		if err != nil {
			t.Error("Unexpected error:", err)
		}

		expected := `{"aps":{"badge":0}}`
		if received := string(bytes); expected != received {
			t.Errorf("Expected %#q but found %#q.", expected, received)
		}
	}

	// Test a badge count of 42.
	{
		payload := NewNotification()
		payload.Badge = 42

		bytes, err := payload.ToJSON()
		if err != nil {
			t.Error("Unexpected error:", err)
		}

		expected := `{"aps":{"badge":42}}`
		if received := string(bytes); expected != received {
			t.Errorf("Expected %#q but found %#q.", expected, received)
		}
	}
}
