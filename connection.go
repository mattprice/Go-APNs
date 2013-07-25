package apns

import (
	"crypto/tls"
	"net"
	"strings"
)

var (
	gatewayConnection *tls.Conn
	connectionCert    tls.Certificate
)

func NewSandboxConnection(cert, key string) (err error) {
	if strings.HasSuffix(cert, ".pem") || strings.HasSuffix(key, ".pem") {
		connectionCert, err = tls.LoadX509KeyPair(cert, key)
	} else {
		connectionCert, err = tls.X509KeyPair([]byte(cert), []byte(key))
	}

	if err != nil {
		return err
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{connectionCert},
	}

	conn, err := net.Dial("tcp", "gateway.sandbox.push.apple.com:2195")
	if err != nil {
		return err
	}

	// TODO: Only one variable for all connections? Doesn't seem smart.
	gatewayConnection = tls.Client(conn, conf)
	err = gatewayConnection.Handshake()
	if err != nil {
		return err
	}

	return nil
}
