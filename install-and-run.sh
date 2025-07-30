#!/bin/bash

ROLE=$1
ADDRESS=$2

if [[ -z "$ROLE" || -z "$ADDRESS" ]]; then
  echo "نحوه استفاده: ./install-and-run.sh client 1.2.3.4:9000"
  exit 1
fi

echo "[+] بررسی نصب بودن Go..."
if ! command -v go &> /dev/null; then
  echo "[!] Go نصب نیست. در حال نصب..."
  sudo apt update
  sudo apt install golang -y
fi

echo "[✓] Go نصب است."
echo "[+] اجرای برنامه..."
go run main.go --role=$ROLE --address=$ADDRESS
