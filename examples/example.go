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

	fmt.Println("Localized Notification, Style A:")
	locNotifcationStyleB()
	fmt.Println()

	// fmt.Println("Localized Notification, Style B:")
	// locNotifcationStyleB()
	// fmt.Println()
}

func notificationStyleA() {
	// Create a simple notification:
	payload := &apns.Notification{
		Alert:       "Hello, World! This is a test.",
		Badge:       42,
		Sound:       "Test.aif",
		LaunchImage: "Default.png",

		Custom: apns.Custom{
			"X-HTTP-Referer": "https://github.com/mattprice/Go-APNs/",
		},
	}

	// Output the payload as JSON, for testing.
	str, _ := payload.ToString()
	fmt.Println(str)
}

func notificationStyleB() {
	// Create a simple notification:
	payload := apns.NewNotification()
	payload.Alert = "Hello, World! This is a test."
	payload.Badge = 42
	payload.Sound = "Test.aif"
	payload.LaunchImage = "Default.png"
	payload.Custom["X-HTTP-Referer"] = "https://github.com/mattprice/Go-APNs"

	// Output the payload as JSON, for testing.
	str, _ := payload.ToString()
	fmt.Println(str)
}

func locNotifcationStyleA() {
	// Create a simple notification:
	payload := &apns.LocNotification{
		Alert:       "APP_GREETING",
		Args:        []string{"John", "Smith"},
		Badge:       42,
		Sound:       "Test.aif",
		LaunchImage: "Default.png",

		Custom: apns.Custom{
			"X-HTTP-Referer": "https://github.com/mattprice/Go-APNs/",
		},
	}

	// Output the payload as JSON, for testing.
	str, _ := payload.ToString()
	fmt.Println(str)
}

func locNotifcationStyleB() {
	// Create a simple notification:
	payload := apns.NewLocNotification()
	payload.Alert = "APP_GREETING"
	payload.Args = []string{"John", "Smith"}
	payload.Badge = 42
	payload.Sound = "Test.aif"
	payload.LaunchImage = "Default.png"
	payload.Custom["X-HTTP-Referer"] = "https://github.com/mattprice/Go-APNs"

	// Output the payload as JSON, for testing.
	str, _ := payload.ToString()
	fmt.Println(str)
}
