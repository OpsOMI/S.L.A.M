//go:build embed
// +build embed

package config

import (
	"log"

	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed embeded.crt
var EmbededTSKCertBinary []byte

//go:embed embeded.yaml
var embeddedClientConfig []byte

func LoadEmbeddedConfig() *Configs {
	var cfg Configs
	if len(embeddedClientConfig) == 0 {
		log.Fatal("embeddedClientConfig is empty, embed files missing")
	}
	if err := yaml.Unmarshal(embeddedClientConfig, &cfg); err != nil {
		log.Fatalf("failed to unmarshal embedded config: %v", err)
	}
	return &cfg
}
