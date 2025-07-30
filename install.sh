#!/bin/bash
set -e

REPO_URL="https://github.com/amirebi9156/Rez-reverse-tunnel.git"
INSTALL_DIR="$HOME/Rez-reverse-tunnel"

sudo apt-get update
sudo apt-get install -y git curl golang

if [ -d "$INSTALL_DIR" ]; then
    echo "[+] Updating repository..."
    git -C "$INSTALL_DIR" pull --ff-only
else
    echo "[+] Cloning repository..."
    git clone "$REPO_URL" "$INSTALL_DIR"
fi

cd "$INSTALL_DIR"

go mod tidy

go build -o tunnel main.go

echo "[+] Launching management console..."
./tunnel
