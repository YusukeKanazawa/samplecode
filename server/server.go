package server

import (
	"log"
	"net"
	"os"
)
type server struct {
}
func (s *server) Listen(){
  // Listen on TCP port 2000 on all interfaces.
  l, err := net.Listen("tcp", ":2000")
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Server Listing on tcp 2000")
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

func New() *server{
  return &server{}
}
