package mqtt

import (
)


type mqtt struct {
	Rcv chan []byte
}

type client struct {
	*mqtt
}

func NewClient() *client {
	return &client{
		&mqtt{
			Rcv: make(chan []byte, 1000),
		},
	}
}

func (c *client) Receive() {
  p := NewParser()
	for {
		buff := <-c.Rcv
    p.Parse(buff)
		//log.Println(string(buff))
	}
}
