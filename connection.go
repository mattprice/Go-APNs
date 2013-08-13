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
	productionConnection *gatewayConnection
	productionConfig     *tls.Config
	sandboxConnection    *gatewayConnection
	sandboxConfig        *tls.Config
)

type gatewayConnection struct {
	client  *tls.Conn
	sandbox bool
}

func LoadCertificates(cert, key string) error {
	keyPair, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return err
	}

	productionConfig = &tls.Config{
		Certificates: []tls.Certificate{keyPair},
	}

	productionConnection = &gatewayConnection{sandbox: false}
	if err := productionConnection.connect(); err != nil {
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

	sandboxConnection = &gatewayConnection{sandbox: true}
	if err := sandboxConnection.connect(); err != nil {
		return err
	}

	return nil
}

func (this *gatewayConnection) connect() error {
	var gateway string
	var config *tls.Config

	if this.sandbox {
		gateway = SANDBOX_GATEWAY
		config = sandboxConfig
	} else {
		gateway = PRODUCTION_GATEWAY
		config = productionConfig
	}

	conn, err := net.Dial("tcp", gateway)
	if err != nil {
		return err
	}

	this.client = tls.Client(conn, config)
	err = this.client.Handshake()
	if err != nil {
		return err
	}

	return nil
}
