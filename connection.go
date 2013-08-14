package apns

import (
	"crypto/tls"
	"net"
	"time"
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
	config  *tls.Config
	gateway string
}

func LoadCertificates(cert, key string) error {
	keyPair, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return err
	}

	productionConfig = &tls.Config{
		Certificates: []tls.Certificate{keyPair},
	}

	productionConnection = &gatewayConnection{
		gateway: PRODUCTION_GATEWAY,
		config:  productionConfig,
	}
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

	sandboxConnection = &gatewayConnection{
		gateway: SANDBOX_GATEWAY,
		config:  sandboxConfig,
	}
	if err := sandboxConnection.connect(); err != nil {
		return err
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

func (this *gatewayConnection) Write(payload []byte) (int, error) {
	return this.client.Write(payload)
}

func (this *gatewayConnection) ReadErrors() []byte {
	_ = this.client.SetReadDeadline(time.Now().Add(5 * time.Second))

	buffer := make([]byte, 6, 6)
	this.client.Read(buffer)

	return buffer
}
