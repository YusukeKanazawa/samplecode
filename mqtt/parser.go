package mqtt

import (
	"log"
  "./message"
)

type parsState byte

const (
	STATE_HEAD parsState = iota + 1
	STATE_VAR_HEAD
	STATE_PAYLOAD
)

type parser struct {
	state parsState
}

func (p *parser) Parse(buff []byte) {
  cursor := 0
	log.Printf("state: %d\n", p.state)

	switch p.state {
	case STATE_HEAD:
    msg := message.New(buff[cursor++]) 
    for ;;{
      
    }
	default:
	}
}

func NewParser() *parser {
	return &parser{
		state: STATE_HEAD,
	}
}
