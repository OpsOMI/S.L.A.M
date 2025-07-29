package models

type Server struct {
	Host                 string `yaml:"host"`
	Port                 string `yaml:"port"`
	RequireAuth          string `yaml:"require_auth"`
	MessageLifeTimeHours int    `yaml:"message_lifetime_hours"`
	LogLevel             string `yaml:"log_level"`
	MaxClients           int    `yaml:"max_clients"`
	TSLCertPath          string `yaml:"tls_cert_path"`
	TSLKeyPath           string `yaml:"tls_key_path"`
}
