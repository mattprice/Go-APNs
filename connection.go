package apns

import (
	"crypto/tls"
	"net"
	"strings"
)

const (
	productionAddress = "gateway.push.apple.com:2195"
	sandboxAddress    = "gateway.sandbox.push.apple.com:2195"
)

var (
	gatewayConnection *tls.Conn
	gatewayAddress    string
	connectionCert    tls.Certificate
)

func NewConnection(cert, key string) error {
	gatewayAddress = productionAddress
	return connection(cert, key)
}

func NewSandboxConnection(cert, key string) error {
	gatewayAddress = sandboxAddress
	return connection(cert, key)
}

func connection(cert, key string) (err error) {
	if strings.HasSuffix(cert, ".pem") || strings.HasSuffix(key, ".pem") {
		// Load the certificate from files.
		connectionCert, err = tls.LoadX509KeyPair(cert, key)
	} else {
		// Load the certificate from input strings.
		connectionCert, err = tls.X509KeyPair([]byte(cert), []byte(key))
	}
	if err != nil {
		return err
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{connectionCert},
	}

	conn, err := net.Dial("tcp", gatewayAddress)
	if err != nil {
		return err
	}

	// TODO: Connection information probably needs to be handled better to support
	// creating production and sandbox connections at the same time.
	gatewayConnection = tls.Client(conn, conf)
	err = gatewayConnection.Handshake()
	if err != nil {
		return err
	}

	return nil
}
