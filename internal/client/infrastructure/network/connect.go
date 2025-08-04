package network

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"time"
)

// ConnectToServer tries to establish a TLS connection to the server with retries.
// serverHost: server IP or domain name
// serverPort: server port number
// certPath: path to the server's root CA certificate file
// timeoutSec: connection timeout in seconds
// retryCount: number of retry attempts if connection fails
func ConnectToServer(
	serverName, serverHost, serverPort, certPath string,
	timeoutSec, retryCount int,
) (net.Conn, error) {
	addr := fmt.Sprintf("%s:%s", serverHost, serverPort)

	// Load the root CA certificate
	certPool := x509.NewCertPool()
	caCert, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, fmt.Errorf("failed to append certificate to pool")
	}

	tlsConfig := &tls.Config{
		RootCAs:    certPool,
		ServerName: serverName,
		MinVersion: tls.VersionTLS12,
	}

	var conn net.Conn
	for i := 0; i <= retryCount; i++ {
		dialer := &net.Dialer{
			Timeout: time.Duration(timeoutSec) * time.Second,
		}
		conn, err = tls.DialWithDialer(dialer, "tcp", addr, tlsConfig)
		if err == nil {
			return conn, nil // successful connection
		}
		if i < retryCount {
			time.Sleep(1 * time.Second) // wait before retrying
		}
	}

	return nil, fmt.Errorf("failed to connect after %d retries: %w", retryCount, err)
}
