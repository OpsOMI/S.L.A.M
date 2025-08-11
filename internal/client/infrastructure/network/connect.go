package network

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/client/config"
)

// ConnectToServer tries to establish a TLS connection to the server with retries.
// serverHost: server IP or domain name
// serverPort: server port number
// certPath: path to the server's root CA certificate file
// timeoutSec: connection timeout in seconds
// retryCount: number of retry attempts if connection fails
func ConnectTsoServer(
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

func ConnectToServer(
	serverName, serverHost, serverPort string,
	certData []byte,
	timeoutSec, retryCount int,
) (net.Conn, error) {
	addr := fmt.Sprintf("%s:%s", serverHost, serverPort)

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(certData); !ok {
		return nil, fmt.Errorf("failed to append certificate to pool")
	}

	tlsConfig := &tls.Config{
		RootCAs:    certPool,
		ServerName: serverName,
		MinVersion: tls.VersionTLS12,
	}

	var conn net.Conn
	var err error
	for i := 0; i <= retryCount; i++ {
		dialer := &net.Dialer{Timeout: time.Duration(timeoutSec) * time.Second}
		conn, err = tls.DialWithDialer(dialer, "tcp", addr, tlsConfig)
		if err == nil {
			return conn, nil
		}
		if i < retryCount {
			time.Sleep(1 * time.Second)
		}
	}

	return nil, fmt.Errorf("failed to connect after %d retries: %w", retryCount, err)
}

func Reconnect(
	conn net.Conn,
	cfg config.Configs,
) (net.Conn, error) {
	if conn != nil {
		conn.Close()
	}

	var certData []byte
	var err error
	if strings.EqualFold(cfg.UseEmbed, "true") {
		certData = config.EmbededTSKCertBinary
	} else {
		certData, err = os.ReadFile(cfg.TSLCertPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read certificate file: %w", err)
		}
	}

	newConn, err := ConnectToServer(
		cfg.ServerName,
		cfg.ServerHost,
		cfg.ServerPort,
		certData,
		cfg.TimeoutSeconds,
		cfg.ReconnectRetry,
	)
	if err != nil {
		return nil, err
	}

	return newConn, nil
}
