package handler

import (
	"fmt"
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
	OnWriteError(conn *Conn, err error)
}

func NewTCPServer(handler TCPServerHandler, config *TCPServerConfig) *TCPServer {
	s := &TCPServer{
		ConnMap: ds.NewMap[string, *Conn](1024),
		handler: handler,
		config:  config,
	}
	return s
}

func (s *TCPServer) Run() error {
	var err error
	s.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return err
	}

	for {
		c, err := s.listener.Accept()
		if err != nil {
			if s.closed.Load() {
				return nil
			}
			return err
		}

		go func(c net.Conn) {
			conn := newConn(c, s.handler, s.config.ConnConfig)

			s.ConnMap.Store(conn.Id, conn)
			s.handler.OnOpen(conn)
			conn.run()
			s.handler.OnClose(conn)
			s.ConnMap.Delete(conn.Id)
		}(c)
	}
}

func (s *TCPServer) Close() {
	if !s.closed.CompareAndSwap(false, true) {
		return
	}

	s.listener.Close()
}
