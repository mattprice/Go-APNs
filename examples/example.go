// Example is used for testing the JSON output of the Go-APNs package during development.
package main

import (
	"fmt"
	"github.com/mattprice/Go-APNs"
)

func main() {
	fmt.Println("Regular Notification, Style A:")
	notificationStyleA()
	fmt.Println()

	// fmt.Println("Regular Notification, Style B:")
	// notificationStyleB()
	// fmt.Println()

	// fmt.Println("Localized Notification, Style A:")
	// notificationLocStyleA()
	// fmt.Println()

	// fmt.Println("Localized Notification, Style B:")
	// notificationLocStyleB()
	// fmt.Println()
}

func notificationStyleA() {
	// Create a simple notification:
	payload := &apns.Notification{
		Alert:       "Hello, World! This is a test.",
		Badge:       42,
		Sound:       "Test.aiff",
		LaunchImage: "Default.png",

		Custom: apns.Custom{
			"X-HTTP-Referer": "https://github.com/mattprice/Go-APNs/",
		},
	}

	// Output the payload as JSON, for testing.
	str, err := payload.ToString()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(str)
	}
}

func notificationStyleB() {
	// Create a simple notification:
	payload := apns.NewNotification()
	payload.Alert = "Hello, World! This is a test."
	payload.Badge = 42
	payload.Sound = "Test.aiff"
	payload.LaunchImage = "Default.png"
	payload.Custom["X-HTTP-Referer"] = "https://github.com/mattprice/Go-APNs"

	// Output the payload as JSON, for testing.
	str, _ := payload.ToString()
	fmt.Println(str)
}

func notificationLocStyleA() {
	// Create a simple notification:
	payload := &apns.Notification{
		Alert:        "Test! This should be ignored.",
		AlertLocKey:  "APP_GREETING",
		AlertLocArgs: []string{"John", "Smith"},
		Badge:        42,
		Sound:        "Test.aiff",
		LaunchImage:  "Default.png",

		Custom: apns.Custom{
			"X-HTTP-Referer": "https://github.com/mattprice/Go-APNs/",
		},
	}

	// Output the payload as JSON, for testing.
	str, _ := payload.ToString()
	fmt.Println(str)
}

func notificationLocStyleB() {
	// Create a simple notification:
	payload := apns.NewNotification()
	payload.Alert = "Test! This should be ignored."
	payload.AlertLocKey = "APP_GREETING"
	payload.AlertLocArgs = []string{"John", "Smith"}
	payload.Badge = 42
	payload.Sound = "Test.aiff"
	payload.LaunchImage = "Default.png"
	payload.Custom["X-HTTP-Referer"] = "https://github.com/mattprice/Go-APNs"

	// Output the payload as JSON, for testing.
	str, _ := payload.ToString()
	fmt.Println(str)
}
