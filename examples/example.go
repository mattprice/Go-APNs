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

	// Add a multilevel map to the payload.
	// You can also send an array, if that's your thing.
	searchEngines := map[string]string{
		"ddg":    "http://duckduckgo.com/",
		"bing":   "http://bing.com/",
		"google": "http://google.com/",
	}
	payload.SetCustom("searchEngines", searchEngines)

	// Add a custom string to the payload.
	payload.SetCustom("defaultEngine", "ddg")

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
