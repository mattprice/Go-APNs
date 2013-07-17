// Example is used for testing the JSON output of the Go-APNs package during development.
package main

import (
	"fmt"
	"github.com/mattprice/Go-APNs"
)

func main() {
	// Create a simple notification:
	payload := &apns.Notification{
		Alert: "Hello, World! This is a test.",
	}

	bytes, err := payload.ToBytes()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Bytes:", bytes)
}
