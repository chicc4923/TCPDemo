package main

import (
	"TCPDemo/tnet"
)

func main() {
	s := tnet.NewServer("DemoV0.1")
	s.ServerRun()
}
