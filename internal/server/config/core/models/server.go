package models

type Server struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	TSLCertPath string `yaml:"tls_cert_path"`
	TSLKeyPath  string `yaml:"tls_key_path"`
}
