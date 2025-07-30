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

# Build binaries if not already built
if [ ! -f server.bin ]; then
    echo "[+] Building server binary..."
    go build -o server.bin ./server
fi

if [ ! -f client.bin ]; then
    echo "[+] Building client binary..."
    go build -o client.bin ./client
fi

case $choice in
  1)
    echo "[+] Starting in SERVER mode (Iran VPS)..."
    ./server.bin
    ;;
  2)
    echo -n "Enter server IP:PORT to connect to: "
    read serverAddr
    echo -n "Enter token: "
    read token
    ./client.bin $serverAddr $token
    ;;
  3)
    echo "Exiting..."
    exit 0
    ;;
  *)
    echo "Invalid option"
    ;;
esac
