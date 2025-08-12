package network

import (
	"crypto/tls"
	"net"
)

func StartServer(
	mode, port, certFile, keyFile string,
) (net.Listener, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	listener, err := tls.Listen("tcp", ":"+port, config)
	if err != nil {
		return nil, err
	}

	return listener, err
}
