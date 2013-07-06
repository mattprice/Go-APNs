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
	payload := &apns.Notification{
		Alert:       "Hello, World! This is a test.",
		Badge:       42,
		Sound:       "Test.aif",
		LaunchImage: "Default.png",
	}

	// TODO: This really doesn't feel as elegant as I'd like.
	payload.SetCustom("link", "http://www.google.com/")

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
