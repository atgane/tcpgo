package handler

import "net"

type Conn[T any] struct {
	Field T // 사용자 정의 field
	conn  net.Conn
	h     ConnHandler[T]
}

type ConnConfig struct {
	// config field 추가
}

func NewConnConfig() *ConnConfig {
	c := &ConnConfig{}
	// config field 초기화
	return c
}

func NewConn[T any](conn net.Conn, h ConnHandler[T], config *ConnConfig) *Conn[T] {
	c := &Conn[T]{}
	c.conn = conn
	c.h = h
	// config에 대한 처리
	return c
}

// connection의 동작 구현
func (c *Conn[T]) Run() error {
	return nil
}

// connection 닫힘 구현
func (c *Conn[T]) Close() {
}
