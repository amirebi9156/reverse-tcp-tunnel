package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
	"github.com/pelletier/go-toml"
)

type ClientConfig struct {
	TunnelName string `toml:"tunnel_name"`
	ServerAddr string `toml:"server_address"`
	Token      string `toml:"token"`
	LocalPort  string `toml:"local_port"`
}

func loadClientConfig(path string) (ClientConfig, error) {
	var config ClientConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = toml.Unmarshal(data, &config)
	return config, err
}

func Start() error {
	config, err := loadClientConfig("config.toml")
	if err != nil {
		return fmt.Errorf("[!] Failed to load config: %v", err)
	}

	fmt.Println("[+] Connecting to tunnel server:", config.ServerAddr)
	conn, err := net.Dial("tcp", config.ServerAddr)
	if err != nil {
		return fmt.Errorf("[!] Failed to connect: %v", err)
	}
	defer conn.Close()

	message := fmt.Sprintf("TOKEN:%s\n", config.Token)
	_, err = conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("[!] Error sending token: %v", err)
	}

	heartbeatInterval := 30 * time.Second
	go func() {
		ticker := time.NewTicker(heartbeatInterval)
		defer ticker.Stop()
		for range ticker.C {
			conn.Write([]byte("HEARTBEAT\n"))
		}
	}()

	ln, err := net.Listen("tcp", ":"+config.LocalPort)
	if err != nil {
		return fmt.Errorf("[!] Failed to listen on local port: %v", err)
	}
	defer ln.Close()

	fmt.Println("[✓] Tunnel ready. Forwarding local port", config.LocalPort)

	for {
		localConn, err := ln.Accept()
		if err != nil {
			fmt.Println("[!] Accept error:", err)
			continue
		}
		go proxy(localConn, conn)
	}
}

func proxy(localConn net.Conn, remoteConn net.Conn) {
	defer localConn.Close()
	go io.Copy(remoteConn, localConn)
	io.Copy(localConn, remoteConn)
}