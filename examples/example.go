// Example is used for testing the JSON output of the Go-APNs package during development.
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/mattprice/Go-APNs"
	"time"
)

func main() {
	// Attempt connection.
	err := apns.NewSandboxConnection("certs/Cert.pem", "certs/Key.pem")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Create a simple notification:
	payload := &apns.Notification{
		Alert: "Hello, World! This is test three!",
	}
	payload.SetExpiryDuration(24 * time.Hour)
	payload.SendTo("41e89803e31c5becfe3bdf2a1862cdd4e60112c0740dae61fe9912fc4eb64d43")

	doDebug(payload)
}

func doDebug(payload *apns.Notification) {
	byteToken, err := hex.DecodeString("41e89803e31c5becfe3bdf2a1862cdd4e60112c0740dae61fe9912fc4eb64d43")
	if err != nil {
		fmt.Println("Error:", err)
	}

	output, err := payload.ToBytes(byteToken)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Test the Bytes output.
	// TODO: Move this into its own function.
	var expiry uint32
	buffer := bytes.NewBuffer(output[5:9])
	binary.Read(buffer, binary.BigEndian, &expiry)

	fmt.Println("Binary Output:")
	fmt.Println("- Command:\t", output[0])
	fmt.Println("- Identifier:\t", output[1:5])
	fmt.Println("- Expiry:\t", expiry)
	fmt.Println("- Token Len:\t", output[9:11])
	fmt.Println("- Token:\t", hex.EncodeToString(output[11:43]))
	fmt.Println("- Paylod Len:\t", output[43:45])
	fmt.Println("- Payload:\t", string(output[45:]))
}
