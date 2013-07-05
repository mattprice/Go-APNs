// Example is used for testing the JSON output of the Go-APNs package during development.
package main

import (
	"fmt"
	"github.com/mattprice/Go-APNs"
)

func main() {
	fmt.Println("Regular Notification:")
	notification()
	fmt.Println()

	// fmt.Println("Localized Notification:")
	// locNotifcation()
	// fmt.Println()
}

func notification() {
	// Create a simple notification:
	payload := &apns.Notification{
		Alert:       "Hello, World! This is a test.",
		Badge:       42,
		Sound:       "Test.aif",
		LaunchImage: "Default.png",
	}

	// Add a custom string to the payload.
	payload.SetCustom("google", "http://www.google.com/")

	// Add a multi-level dictionary to the payload.
	multipleLinks := map[string]string{
		"ddg":  "http://duckduckgo.com/",
		"bing": "http://bing.com/",
	}
	payload.SetCustom("links", multipleLinks)

	// Add an array to the payload.
	arrayOfNames := []string{"John", "Sue"}
	payload.SetCustom("names", arrayOfNames)

	// Output the payload as JSON, for testing.
	str, _ := payload.ToString()
	fmt.Println(str)
}

// func locNotifcation() {
// 	payload := &apns.LocNotification{
// 		Alert: "Hello, World! This is also a test!",
// 		Badge: 42,
// 	}

// 	str, _ := payload.ToString()
// 	fmt.Println(str)
// }
