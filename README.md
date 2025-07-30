# Reverse TCP Tunnel

This project is a simple reverse TCP tunneling tool written in Go. It allows a client running on a foreign VPS to forward connections back to an Iranian server. The configuration is stored in `config.toml`.

## Features

- JSON based handshake with token authentication
- Multiple tunnel ports
- Configurable heartbeat interval
- Logging to console and optional file

## Usage

Run the installer directly from GitHub:

```bash
bash <(curl -s https://raw.githubusercontent.com/amirebi9156/Rez-reverse-tunnel/main/install.sh --ipv4)
```

The script installs Go if needed, fetches Go modules, builds the binaries,
and then presents an interactive menu where you can start the server or client.
