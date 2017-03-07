package main

import (
	"../server"
)

func main() {
	svr := server.New(1883)
	svr.Listen()
}
