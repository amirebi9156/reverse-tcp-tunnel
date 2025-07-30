# Reverse TCP Tunnel

This project is a reverse TCP tunneling tool written in Go. It forwards connections from a client VPS back to a server VPS. Configuration is stored in `config.toml` which can be generated interactively.

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

The script installs Go and Git if needed, clones this repository, builds the binaries and then launches the interactive management console.

### نحوه استفاده (فارسی)

```bash
bash <(curl -s https://raw.githubusercontent.com/amirebi9156/Rez-reverse-tunnel/main/install.sh --ipv4)
```

پس از اجرای دستور بالا، اسکریپت وابستگی‌ها را نصب کرده و برنامه را اجرا می‌کند. با انتخاب Server یا Client تنظیمات موردنیاز پرسیده می‌شود و فایل `config.toml` ساخته خواهد شد.
