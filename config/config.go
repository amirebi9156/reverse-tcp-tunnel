// config/config.go

package config

import (
	"log"

	"github.com/pelletier/go-toml"
)

type Config struct {
	ListenAddr  string `toml:"listen_addr"`
	ConnectAddr string `toml:"connect_addr"`
}

func Load(path string) Config {
	var cfg Config
	tree, err := toml.LoadFile(path)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = tree.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	return cfg
}
