package models

type DatabaseConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
	Timezone string
}
