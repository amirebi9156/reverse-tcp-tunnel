package server

import (
	"bufio"
	"encoding/json"
	"io"
	"net"
	"time"

	"reverse/config"
	"reverse/pkg/logger"
)

type handshake struct {
	Name  string `json:"name"`
	Token string `json:"token"`
	Port  string `json:"port"`
}

func Start(cfg config.Config) error {
	if err := logger.Init(cfg.LogFile); err != nil {
		return err
	}

	ln, err := net.Listen("tcp", cfg.ListenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	logger.Log.Printf("server listening on %s", cfg.ListenAddr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Log.Printf("accept error: %v", err)
			continue
		}
		go handleConnection(conn, cfg)
	}
}

func handleConnection(conn net.Conn, cfg config.Config) {
	logger.Log.Printf("connection from %s", conn.RemoteAddr())
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		logger.Log.Printf("handshake read error: %v", err)
		return
	}
	conn.SetReadDeadline(time.Time{})

	var hs handshake
	if err := json.Unmarshal([]byte(line), &hs); err != nil {
		logger.Log.Printf("invalid handshake: %v", err)
		return
	}
	if hs.Token != cfg.Token {
		logger.Log.Printf("invalid token")
		return
	}

	logger.Log.Printf("tunnel %s requested on port %s", hs.Name, hs.Port)

	if !contains(cfg.TunnelPorts, hs.Port) {
		logger.Log.Printf("unauthorized port %s", hs.Port)
		return
	}

	target := "127.0.0.1:" + hs.Port
	targetConn, err := net.Dial("tcp", target)
	if err != nil {
		logger.Log.Printf("failed to connect to %s: %v", target, err)
		return
	}
	defer targetConn.Close()

	logger.Log.Printf("tunneling to %s", target)
	go io.Copy(targetConn, reader)
	io.Copy(conn, targetConn)
}

func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}
