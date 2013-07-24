package apns

import (
	"crypto/tls"
	"net"
	"strings"
)

func NewSandboxConnection(cert, key string) (err error) {
	var connCert tls.Certificate

	if strings.HasSuffix(cert, ".pem") || strings.HasSuffix(key, ".pem") {
		connCert, err = tls.LoadX509KeyPair(cert, key)
	} else {
		connCert, err = tls.X509KeyPair([]byte(cert), []byte(key))
	}

	if err != nil {
		return err
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{connCert},
	}

	conn, err := net.Dial("tcp", "gateway.sandbox.push.apple.com:2195")
	if err != nil {
		return err
	}

	tlsConn := tls.Client(conn, conf)
	err = tlsConn.Handshake()
	if err != nil {
		return err
	}

	return nil
}
