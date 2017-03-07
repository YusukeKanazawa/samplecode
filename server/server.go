package server

import (
	"log"
	"net"
	"os"
	"strconv"
)
type server struct {
	port int
}
func (s *server) Listen(){
  // Listen on TCP port 2000 on all interfaces.
  l, err := net.Listen("tcp", ":" + strconv.Itoa(s.port))
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Server Listing on tcp 2000")

  //broker := NewBroker()

  defer l.Close()
  for {
    // Wait for a connection.
    conn, err := l.Accept()
    if err != nil {
      log.Fatal(err)
      os.Exit(1)
    }
    go handleClient(conn)
  }
}

func New(port int) *server{
  return &server{
		port: port,
	}
}
