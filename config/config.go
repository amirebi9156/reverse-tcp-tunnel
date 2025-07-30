// config/config.go

package config

import (
	"bytes"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Name        string   `toml:"name"`
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

// Save writes the configuration to the specified file in TOML format.
func Save(path string, cfg Config) error {
	buf := &bytes.Buffer{}
	if err := toml.NewEncoder(buf).Encode(cfg); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0644)
}
