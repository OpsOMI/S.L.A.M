package models

type Server struct {
	ExternalHost string `yaml:"external_host"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	TSLCertPath  string `yaml:"tls_cert_path"`
	TSLKeyPath   string `yaml:"tls_key_path"`
}
