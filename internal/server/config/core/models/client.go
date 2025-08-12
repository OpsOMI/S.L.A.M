package models

type ClientConfig struct {
	ClientKey      string `yaml:"client_key"`
	ServerHost     string `yaml:"server_host"`
	ServerPort     string `yaml:"server_port"`
	TSLServerName  string `yaml:"tsl_server_name"`
	TimeoutSeconds int    `yaml:"timeout_seconds"`
	ReconnectRetry int    `yaml:"reconnect_retry"`
}
