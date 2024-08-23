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

func (c *Conn) Write(b []byte) error {
	var err error
	p := 0
	write := 0
	l := len(b)
	for write < l {
		p, err = c.conn.Write(b[write:])
		if err != nil {
			if c.closed.Load() {
				return nil
			}
			return err
		}
		write += p
	}
	return nil
}

func (c *Conn) Close() {
	if !c.closed.CompareAndSwap(false, true) {
		return
	}

	c.conn.Close()
}

func (c *Conn) run() {
	b := make([]byte, c.config.BufferSize)
	buf := make([]byte, c.config.BufferSize)
	for {
		n, err := c.conn.Read(b)
		if err != nil {
			if c.closed.Load() {
				return
			}
			c.handler.OnReadError(c, err)
			return
		}

		buf = append(buf, b[:n]...)
		p := 0
		read := 0
		l := len(buf)
		for read < l {
			p = c.handler.OnRead(c, buf[read:])
			if p == 0 {
				break
			}
			read += p
		}
		buf = buf[read:]
	}
}
