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
	// Create a simple notification:
	payload := &apns.Notification{
		Alert: "Hello, World! This is a test.",
	}

	payload.SetExpiryTime(time.Now().Add(24 * time.Hour))

	output, err := payload.ToBytes()
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
