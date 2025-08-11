package models

type Server struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	TSLServerName string `yaml:"tsl_server_name"`
	TSLCertPath   string `yaml:"tls_cert_path"`
	TSLKeyPath    string `yaml:"tls_key_path"`
}
