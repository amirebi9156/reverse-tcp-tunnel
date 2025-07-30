// main.go — Reverse TCP Tunnel with Configurable Port and Interactive Menu in Go

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"reverse/client"
	"reverse/config"
	"reverse/server"
)

func main() {
	_, err := config.Load("config.toml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
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
		fmt.Println("[+] Starting in SERVER mode (Iran VPS)...")
		if err := server.Start(); err != nil {
			log.Fatal("[!] Server error:", err)
		}
	case "2":
		fmt.Println("[+] Starting in CLIENT mode (Foreign VPS)...")
		if err := client.Start(); err != nil {
			log.Fatal("[!] Client error:", err)
		}
	case "8":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please select a number between 1 and 8.")
		os.Exit(1)
	}
}
