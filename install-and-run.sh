#!/bin/bash

clear
echo "  **** utunnel Reverse Tunnel Management Console ****"
echo ""
echo "  1) Start in Server Mode"
echo "  2) Start in Client Mode"
echo "  3) Exit"
echo ""
echo -n "Enter your choice (1-3): "
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
  3)
    echo "Exiting..."
    exit 0
    ;;
  *)
    echo "Invalid option"
    ;;
esac
