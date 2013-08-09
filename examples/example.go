// Example is used for testing the JSON output of the Go-APNs package during development.
package main

import (
	"fmt"
	"github.com/mattprice/Go-APNs"
	"time"
)

func main() {
	// Attempt connection.
	err := apns.LoadSandboxCertificates("certs/Cert.pem", "certs/Key.pem")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Create a simple notification:
	payload := &apns.Notification{
		Alert: "Hello, World! This is a test!",
	}
	payload.SetExpiryDuration(24 * time.Hour)
	payload.DebugBinary("41e89803e31c5becfe3bdf2a1862cdd4e60112c0740dae61fe9912fc4eb64d43")
	payload.SendTo("41e89803e31c5becfe3bdf2a1862cdd4e60112c0740dae61fe9912fc4eb64d43")
}
