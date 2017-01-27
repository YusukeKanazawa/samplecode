package server

import (
	"../mqtt"
	"log"
)

type session struct {
	rcv chan []byte
}

func newSession() *session {
	return &session{
		rcv: make(chan []byte, 100),
	}
}

func (s *session) receive() {
	for {
		b := <-s.rcv
		msg := mqtt.NewMessage(b)
	}
}
