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
  KEEP_ALIVE = 60
)

func handleClient(c net.Conn) {
	log.Printf("Connected. IP:%s\n", c.LocalAddr())
  
  client := mqtt.NewClient()
  go client.Receive()

	defer c.Close()
	buff := make([]byte, BUFFER_SIZE)
	for {
		  c.SetReadDeadline(time.Now().Add(KEEP_ALIVE * time.Second))
  		l, err := c.Read(buff)

  		if err != nil {
  			log.Fatal(err.Error())
  			os.Exit(1)
  		}
      client.Rcv <- buff[:l]
      
		//c.SetWriteDeadline(time.Now().Add(KEEP_ALIVE * time.Second))
	}
}
