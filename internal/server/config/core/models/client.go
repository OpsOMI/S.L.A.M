package models

type ClientConfig struct {
	ClientID       string `yaml:"client_id"`
	ServerName     string `yaml:"server_name"`
	ServerHost     string `yaml:"server_host"`
	ServerPort     string `yaml:"server_port"`
	TSLCertPath    string `yaml:"tls_cert_path"`
	TimeoutSeconds int    `yaml:"timeout_seconds"`
	ReconnectRetry int    `yaml:"reconnect_retry"`
}
