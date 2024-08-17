package main

import (
	"net"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	listener, err := net.Listen("tcp", ":1212")
	if err != nil {
		log.Fatal().Err(err).Msg("tcp listen error")
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Fatal().Err(err).Msg("tcp accept error")
		}

		go func() {
			b := make([]byte, 1024)
			for {
				n, err := c.Read(b)
				if err != nil {
					log.Error().Err(err).Msg("connection read error")
					return
				}

				log.Trace().Int("readSize", n).Bytes("data", b[:n]).Msg("connection read success")

				_, err = c.Write(b[:n])
				if err != nil {
					log.Error().Err(err).Msg("connection write error")
					return
				}
			}
		}()
	}
}
