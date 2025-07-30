package client

import (
	"encoding/json"
	"io"
	"net"
	"time"

	"reverse/config"
	"reverse/pkg/logger"
)

// handshake message structure
type handshake struct {
	Name  string `json:"name"`
	Token string `json:"token"`
	Port  string `json:"port"`
}

func Start(cfg config.Config) error {
	if err := logger.Init(cfg.LogFile); err != nil {
		return err
	}

	for {
		if err := run(cfg); err != nil {
			logger.Log.Printf("connection error: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	return nil
}

func run(cfg config.Config) error {
	logger.Log.Printf("connecting to %s", cfg.ConnectAddr)
	conn, err := net.DialTimeout("tcp", cfg.ConnectAddr, 10*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()

	hs := handshake{Name: cfg.Name, Token: cfg.Token, Port: cfg.TunnelPorts[0]}
	data, _ := json.Marshal(hs)
	conn.Write(append(data, '\n'))

	go heartbeat(conn, cfg.Heartbeat)

	ln, err := net.Listen("tcp", ":"+cfg.TunnelPorts[0])
	if err != nil {
		return err
	}
	defer ln.Close()

	logger.Log.Printf("tunnel ready on local port %s", cfg.TunnelPorts[0])
	for {
		localConn, err := ln.Accept()
		if err != nil {
			logger.Log.Printf("accept error: %v", err)
			continue
		}
		go proxy(localConn, conn)
	}
}

func heartbeat(conn net.Conn, interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		conn.Write([]byte("{}\n"))
	}
}

func proxy(localConn net.Conn, remoteConn net.Conn) {
	defer localConn.Close()
	go io.Copy(remoteConn, localConn)
	io.Copy(localConn, remoteConn)
}
