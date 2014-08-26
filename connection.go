package apns

import (
	"crypto/tls"
    "fmt"
	"log"
	"net"
	"time"
)

const (
	PRODUCTION_GATEWAY = "gateway.push.apple.com"
	SANDBOX_GATEWAY    = "gateway.sandbox.push.apple.com"
)

const (
    GATEWAY_PORT = 2195
)

var (
	productionConnection *gatewayConnection
	productionConfig     *tls.Config
	sandboxConnection    *gatewayConnection
	sandboxConfig        *tls.Config
)

type gatewayConnection struct {
	client  *tls.Conn
	config  *tls.Config
	gateway string
}

// TODO: Wrap log.Print calls so we can easily disable them.

func LoadCertificate(production bool, certContents []byte) error {
	keyPair, err := tls.X509KeyPair(certContents, certContents)
	if err != nil {
		return err
	}

	if err := storeAndConnect(production, keyPair); err != nil {
		return err
	}

	return nil
}

func LoadCertificateFile(production bool, certLocation string) error {
	keyPair, err := tls.LoadX509KeyPair(certLocation, certLocation)
	if err != nil {
		return err
	}

	if err := storeAndConnect(production, keyPair); err != nil {
		return err
	}

	return nil
}

func storeAndConnect(production bool, keyPair tls.Certificate) error {
	if production {
		// Production Connections
		productionConfig = &tls.Config{
			Certificates: []tls.Certificate{keyPair},
			ServerName: PRODUCTION_GATEWAY,
		}

		productionConnection = &gatewayConnection{
			gateway: fmt.Sprintf("%v:%v",PRODUCTION_GATEWAY, GATEWAY_PORT),
			config:  productionConfig,
		}

		if err := productionConnection.connect(); err != nil {
			return err
		}
	} else {
		// Sandbox Connections
		sandboxConfig = &tls.Config{
			Certificates: []tls.Certificate{keyPair},
			ServerName: SANDBOX_GATEWAY,
		}

		sandboxConnection = &gatewayConnection{
			gateway: fmt.Sprintf("%v:%v",SANDBOX_GATEWAY, GATEWAY_PORT),
			config:  sandboxConfig,
		}

		if err := sandboxConnection.connect(); err != nil {
			return err
		}
	}

	return nil
}

func (this *gatewayConnection) connect() error {
	conn, err := net.Dial("tcp", this.gateway)
	if err != nil {
		return err
	}

	this.client = tls.Client(conn, this.config)
	err = this.client.Handshake()
	if err != nil {
		return err
	}

	return nil
}

func (this *gatewayConnection) Write(payload []byte) error {
	_, err := this.client.Write(payload)
	if err != nil {
		// We probably disconnected. Reconnect and resend the message.
		// TODO: Might want to check the actual error returned?
		log.Printf("[APNS] Error writing data to socket: %v", err)
		log.Println("[APNS] *** Server disconnected unexpectedly. ***")
		err := this.connect()
		if err != nil {
			log.Printf("[APNS] Could not reconnect to the server: %v", err)
			return err
		}
		log.Println("[APNS] Reconnected to the server successfully.")

		// TODO: This could cause an endless loop of errors.
		// 		If it's the connection failing, that would be caught above.
		// 		So why don't we add a counter to the payload itself?
		this.Write(payload)
	}

	return nil
}

func (this *gatewayConnection) ReadErrors() (bool, []byte) {
	_ = this.client.SetReadDeadline(time.Now().Add(5 * time.Second))

	buffer := make([]byte, 6, 6)
	n, _ := this.client.Read(buffer)

	// n == 0 if there were no errors.
	if n == 0 {
		// TODO: I think this would get returned even if Read() produces an error.
		return false, nil
	}

	return true, buffer
}
