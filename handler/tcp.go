package handler

import (
	"main/ds"
	"net"
	"sync/atomic"
)

type TCPServer struct {
	ConnMap  *ds.Map[string, *Conn]
	listener net.Listener
	handler  TCPServerHandler
	config   *TCPServerConfig
	closed   atomic.Bool
}

type TCPServerConfig struct {
	*ConnConfig
	Port int
}

func NewTCPServerConfig() *TCPServerConfig {
	return &TCPServerConfig{
		ConnConfig: newConnConfig(),
	}
}

type TCPServerHandler interface {
	OnOpen(conn *Conn)
	OnClose(conn *Conn)
	OnRead(conn *Conn, b []byte) (n int)
	OnReadError(conn *Conn, err error)
}

func NewTCPServer(handler TCPServerHandler, config *TCPServerConfig) *TCPServer {
	s := &TCPServer{
		ConnMap: ds.NewMap[string, *Conn](1024),
		handler: handler,
		config:  config,
	}
	return s
}

func (s *TCPServer) Run() error

func (s *TCPServer) Close()
