// main.go — Reverse TCP Tunnel with Configurable Port and Interactive Menu in Go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"reverse/client"
	"reverse/config"
	"reverse/server"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(`
  **** utunnel Reverse Tunnel Management Console ****

  1) Start in Server Mode
  2) Start in Client Mode
  3) Tunnel Status (inactive)
  4) Restart Service (inactive)
  5) Stop Service (inactive)
  6) Delete Service (inactive)
  7) Reset Timer (inactive)
  8) Exit
`)
	fmt.Print("Enter your choice (1-8): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		cfg := askServerConfig(reader)
		config.Save("config.toml", cfg)
		fmt.Println("[+] Starting in SERVER mode...")
		if err := server.Start(cfg); err != nil {
			fmt.Println("Server error:", err)
		}
	case "2":
		cfg := askClientConfig(reader)
		config.Save("config.toml", cfg)
		fmt.Println("[+] Starting in CLIENT mode...")
		if err := client.Start(cfg); err != nil {
			fmt.Println("Client error:", err)
		}
	case "8":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please select a number between 1 and 8.")
		os.Exit(1)
	}
}

func askServerConfig(r *bufio.Reader) config.Config {
	fmt.Print("Tunnel name: ")
	name, _ := r.ReadString('\n')
	fmt.Print("Listen address [0.0.0.0:9000]: ")
	laddr, _ := r.ReadString('\n')
	fmt.Print("Token: ")
	token, _ := r.ReadString('\n')
	fmt.Print("Forward ports (comma separated) [8080]: ")
	portsStr, _ := r.ReadString('\n')

	ports := parsePorts(strings.TrimSpace(portsStr), "8080")

	return config.Config{
		Name:        strings.TrimSpace(name),
		ListenAddr:  defaultStr(strings.TrimSpace(laddr), "0.0.0.0:9000"),
		Token:       strings.TrimSpace(token),
		TunnelPorts: ports,
		Heartbeat:   30,
		LogFile:     "reverse.log",
	}
}

func askClientConfig(r *bufio.Reader) config.Config {
	fmt.Print("Tunnel name: ")
	name, _ := r.ReadString('\n')
	fmt.Print("Server address [1.2.3.4:9000]: ")
	addr, _ := r.ReadString('\n')
	fmt.Print("Token: ")
	token, _ := r.ReadString('\n')
	fmt.Print("Local service port [8080]: ")
	port, _ := r.ReadString('\n')

	return config.Config{
		Name:        strings.TrimSpace(name),
		ConnectAddr: defaultStr(strings.TrimSpace(addr), "1.2.3.4:9000"),
		Token:       strings.TrimSpace(token),
		TunnelPorts: []string{defaultStr(strings.TrimSpace(port), "8080")},
		Heartbeat:   30,
		LogFile:     "reverse.log",
	}
}

func parsePorts(in, def string) []string {
	if in == "" {
		return []string{def}
	}
	fields := strings.Split(in, ",")
	var out []string
	for _, f := range fields {
		f = strings.TrimSpace(f)
		if f != "" {
			out = append(out, f)
		}
	}
	return out
}

func defaultStr(v, def string) string {
	if v == "" {
		return def
	}
	return v
}
