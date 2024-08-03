package main

import (
	"net"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	l, err := net.Listen("tcp", ":3000") // 서버측 소켓 개방
	if err != nil {
		log.Fatal().Err(err).Msg("tcp listen error")
	}

	for {
		c, err := l.Accept() // 클라이언트측 소켓 연결
		if err != nil {
			log.Error().Err(err).Msg("tcp accept error")
			return
		}

		go func() { // 클라이언트 소켓 연결마다 비동기 함수로 로직 처리
			for {
				var err error
				b := make([]byte, 4096) // 버퍼 생성
				n, err := c.Read(b)     // 블로킹 읽기
				if err != nil {
					log.Error().Err(err).Msg("tcp read error")
					return
				}

				time.Sleep(time.Second * 5)

				_, err = c.Write(b[:n]) // 블로킹 쓰기
				if err != nil {
					log.Error().Err(err).Msg("tcp write error")
					return
				}
			}
		}()
	}
}
