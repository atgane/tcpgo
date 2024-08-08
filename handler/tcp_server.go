package handler

import (
	"tcpgo/ds"
)

type TCPServer[T any] struct {
	ds.Map[string, Conn[T]]
}

// 아래 동작 메서드 구현
type ConnHandler[T any] interface {
	OnOpen(*Conn[T])
	OnClose(*Conn[T])
	OnRead(*Conn[T], []byte) uint
	OnReadError(*Conn[T], error)
	OnWriteError(*Conn[T], error)
}

type TCPServerConfig struct {
	Port       int
	ConnConfig *ConnConfig
}

func NewTCPServerConfig() *TCPServerConfig {
	c := &TCPServerConfig{}
	// config field 초기화
	return c
}

func NewTCPServer[T any](config *TCPServerConfig) (*TCPServer[T], error) {
	s := &TCPServer[T]{}
	// config에 대한 처리
	return s, nil
}

// TCP 서버 동작 구현
func (s *TCPServer[T]) Run() error {
	return nil
}

// TCP 서버 닫힘 구현
func (s *TCPServer[T]) Close() {
}
