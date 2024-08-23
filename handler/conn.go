package handler

import (
	"net"
	"sync/atomic"

	"github.com/google/uuid"
)

type Conn struct {
	Id      string
	conn    net.Conn
	handler TCPServerHandler
	config  *ConnConfig
	closed  atomic.Bool
}

func newConnConfig() *ConnConfig {
	return &ConnConfig{
		BufferSize:         4096,
		ReadTimeoutSecond:  600,
		WriteTimeoutSecond: 600,
	}
}

type ConnConfig struct {
	BufferSize         int
	ReadTimeoutSecond  int
	WriteTimeoutSecond int
}

func newConn(
	conn net.Conn,
	handler TCPServerHandler,
	config *ConnConfig,
) *Conn {
	c := &Conn{
		Id:      uuid.New().String(),
		conn:    conn,
		handler: handler,
		config:  config,
	}
	return c
}

func (c *Conn) Write(b []byte) error

func (c *Conn) Close()

func (c *Conn) run()
