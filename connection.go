package apns

import (
	"crypto/tls"
	"net"
)

const (
	PRODUCTION_GATEWAY = "gateway.push.apple.com:2195"
	SANDBOX_GATEWAY    = "gateway.sandbox.push.apple.com:2195"
)

var (
	productionClient *tls.Conn
	productionConfig *tls.Config

	sandboxClient *tls.Conn
	sandboxConfig *tls.Config
)

func LoadCertificates(cert, key string) error {
	keyPair, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return err
	}

	productionConfig = &tls.Config{
		Certificates: []tls.Certificate{keyPair},
	}

	if err := sandboxConnect(); err != nil {
		return err
	}

	return nil
}

func LoadSandboxCertificates(cert, key string) error {
	keyPair, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return err
	}

	sandboxConfig = &tls.Config{
		Certificates: []tls.Certificate{keyPair},
	}

	if err := sandboxConnect(); err != nil {
		return err
	}

	return nil
}

func sandboxConnect() error {
	conn, err := net.Dial("tcp", SANDBOX_GATEWAY)
	if err != nil {
		return err
	}

	// TODO: Don't hardcode this as a sandbox connection.
	sandboxClient = tls.Client(conn, sandboxConfig)
	err = sandboxClient.Handshake()
	if err != nil {
		return err
	}

	return nil
}
