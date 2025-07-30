#!/bin/bash

clear
cat <<EOF
  **** utunnel Reverse Tunnel Management Console ****

  1) Start in Server Mode
  2) Start in Client Mode
  3) Tunnel Status (inactive)
  4) Restart Service (inactive)
  5) Stop Service (inactive)
  6) Delete Service (inactive)
  7) Reset Timer (inactive)
  8) Exit
EOF
echo -n "Enter your choice (1-8): "
read choice

# Check for Go installation
if ! command -v go &> /dev/null; then
    echo "[!] Golang is not installed. Installing..."
    sudo apt update && sudo apt install -y golang
fi
if ! command -v git &> /dev/null; then
    echo "[!] Git is not installed. Installing..."
    sudo apt update && sudo apt install -y git
fi

# Fetch Go module dependencies
go mod download

# Build binaries if not already built
if [ ! -f server.bin ]; then
    echo "[+] Building server binary..."
    go build -o server.bin ./server
fi

if [ ! -f client.bin ]; then
    echo "[+] Building client binary..."
    go build -o client.bin ./client
fi

if [ ! -f config.toml ]; then
    echo "[!] config.toml not found. Creating..."
    read -p "Listen address [0.0.0.0:9000]: " laddr
    read -p "Connect address [1.2.3.4:9000]: " caddr
    read -p "Token: " token
    laddr=${laddr:-0.0.0.0:9000}
    caddr=${caddr:-1.2.3.4:9000}
    cat > config.toml <<EOF
listen_addr = "$laddr"
connect_addr = "$caddr"
token = "$token"
tunnel_ports = ["8080"]
heartbeat_interval = 30
log_file = "reverse.log"
EOF
fi

case $choice in
  1)
    echo "[+] Starting in SERVER mode (Iran VPS)..."
    ./server.bin
    ;;
  2)
    echo "[+] Starting in CLIENT mode (Foreign VPS)..."
    ./client.bin
    ;;
  3|4|5|6|7)
    echo "Option $choice is not implemented yet."
    ;;
  8)
    echo "Exiting..."
    exit 0
    ;;
  *)
    echo "Invalid option"
    ;;
esac
