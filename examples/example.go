// Example is used for testing the JSON output of the Go-APNs package during development.
package main

import (
	"fmt"
	"github.com/mattprice/Go-APNs"
	"time"
)

func main() {
	// Attempt connection.
	err := apns.LoadCertificateFile(false, "certs/sandbox.pem")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Create a simple notification:
	h, m, s := time.Now().Clock()
	payload := &apns.Notification{
		Alert:   fmt.Sprintf("Hello! It's %v:%v:%v.", h, m, s),
		Sandbox: true,
	}
	payload.SetExpiryDuration(24 * time.Hour)
	payload.SendTo("01b67b3ffc8405c1d9ece77b6e4747b97ecdacb4ce940af1fca260b9a0311d80")

	// payload.DebugBinary("01b67b3ffc8405c1d9ece77b6e4747b97ecdacb4ce940af1fca260b9a0311d80")
}
