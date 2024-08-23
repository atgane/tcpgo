package main

import (
	"main/handler"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	h := &TCPServerHandler{}
	config := handler.NewTCPServerConfig()
	config.Port = 1212
	s := handler.NewTCPServer(h, config)
	s.Run()
}

type TCPServerHandler struct{}

func (t *TCPServerHandler) OnOpen(conn *handler.Conn) {
	log.Info().Str("socketId", conn.Id).Msg("connection open")
}

func (t *TCPServerHandler) OnClose(conn *handler.Conn) {
	log.Info().Str("socketId", conn.Id).Msg("connection close")
}

func (t *TCPServerHandler) OnRead(conn *handler.Conn, b []byte) (n int) {
	log.Info().Str("socketId", conn.Id).Msg("connection read")
	conn.Write(b)
	return len(b)
}

func (t *TCPServerHandler) OnReadError(conn *handler.Conn, err error) {
	log.Error().Str("socketId", conn.Id).Err(err).Msg("connection read error")
}
