// config/config.go

package config

import (
	"github.com/pelletier/go-toml"
)

type Config struct {
	ListenAddr  string   `toml:"listen_addr"`
	ConnectAddr string   `toml:"connect_addr"`
	Token       string   `toml:"token"`
	TunnelPorts []string `toml:"tunnel_ports"`
	Heartbeat   int      `toml:"heartbeat_interval"`
	LogFile     string   `toml:"log_file"`
}

func Load(path string) (Config, error) {
	var cfg Config
	tree, err := toml.LoadFile(path)
	if err != nil {
		return cfg, err
	}

	err = tree.Unmarshal(&cfg)
	return cfg, err
}
