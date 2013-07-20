// Example is used for testing the JSON output of the Go-APNs package during development.
package main

import (
	"fmt"
	"github.com/mattprice/Go-APNs"
	"time"
)

func main() {
	// Create a simple notification:
	payload := &apns.Notification{
		Alert: "Hello, World! This is a test.",
	}

	payload.SetExpiryTime(time.Now().Add(24 * time.Hour))

	bytes, err := payload.ToBytes()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Bytes:", bytes)
}
