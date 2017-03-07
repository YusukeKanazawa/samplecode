package server

import (
	"log"
	"net"
	"os"
	"time"
  "../mqtt"
)

const (
	BUFFER_SIZE = 1024
)

func handleClient(c net.Conn) {
	log.Printf("Connected. IP:%s\n", c.LocalAddr())

  sess := mqtt.NewSession()
  go sess.Receive()
	go sess.Send(c)
	defer c.Close()

	for {
		c.SetReadDeadline(time.Now().Add(time.Duration(float64(sess.KeepAlive) * 1.5) * time.Second))
		buff := make([]byte, BUFFER_SIZE)
		l, err := c.Read(buff)

		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}
		sess.Rcv <- mqtt.NewMessage(buff[:l])
		//c.SetWriteDeadline(time.Now().Add(KEEP_ALIVE * time.Second))
	}
}
