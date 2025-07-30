// server/server.go
package server

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

type TunnelConfig struct {
	Name        string   `toml:"tunnel_name"`
	Port        string   `toml:"listen_port"`
	Token       string   `toml:"token"`
	TunnelPorts []string `toml:"tunnel_ports"`
}

func loadServerConfig(path string) (TunnelConfig, error) {
	var config TunnelConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = toml.Unmarshal(data, &config)
	return config, err
}

func Start() error {
	config, err := loadServerConfig("config.toml")
	if err != nil {
		return fmt.Errorf("[!] Failed to load config: %v", err)
	}

	address := ":" + config.Port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("[!] Failed to listen on %s: %v", address, err)
	}
	defer listener.Close()

	fmt.Println("[+] Tunnel server is listening on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[!] Accept error:", err)
			continue
		}
		go handleConnection(conn, config)
	}
}

func handleConnection(conn net.Conn, config TunnelConfig) {
	fmt.Println("[+] Connection from", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("[!] Failed to read token:", err)
		conn.Close()
		return
	}
	tokenLine := strings.TrimSpace(line)
	if tokenLine != "TOKEN:"+config.Token {
		fmt.Println("[!] Invalid token. Closing connection")
		conn.Close()
		return
	}

	fmt.Println("[✓] Valid client authenticated")

	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("[!] Client disconnected or error:", err)
				return
			}
			if strings.TrimSpace(line) == "HEARTBEAT" {
				fmt.Println("[~] Received heartbeat from client")
			}
		}
	}()

	if len(config.TunnelPorts) == 0 {
		fmt.Println("[!] No tunnel ports specified")
		conn.Close()
		return
	}

	target := "127.0.0.1:" + config.TunnelPorts[0]
	targetConn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Println("[!] Failed to connect to local service on", target, "-", err)
		conn.Close()
		return
	}

	go io.Copy(targetConn, conn)
	go io.Copy(conn, targetConn)
}
