# Reverse TCP Tunnel

This project is a simple reverse TCP tunneling tool written in Go. It allows a client running on a foreign VPS to forward connections back to an Iranian server. The configuration is stored in `config.toml`.

## Features

- JSON based handshake with token authentication
- Multiple tunnel ports
- Configurable heartbeat interval
- Logging to console and optional file

## Usage

Build and run using the helper script:

```bash
./install-and-run.sh
```

The script will build the binaries and start either the server or client mode.
