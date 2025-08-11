//go:build !embed
// +build !embed

package config

import (
	"log"
)

var EmbededTSKCertBinary []byte

func LoadEmbeddedConfig() *Configs {
	EmbededTSKCertBinary = nil
	log.Fatal("LoadEmbeddedConfig called but no embed build tag set")
	return nil
}
