package main

import (
	"../server"
)

func main() {
	svr := server.New()
	svr.Listen()
}
