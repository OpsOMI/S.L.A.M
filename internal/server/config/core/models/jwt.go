package models

type Jwt struct {
	Issuer string `yaml:"issuer"`
	Secret string `yaml:"secret"`
}
