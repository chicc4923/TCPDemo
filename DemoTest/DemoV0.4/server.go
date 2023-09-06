package main

import (
	"TCPDemo/tface"
	"TCPDemo/tnet"
	"fmt"
)

type PingRouter struct {
	tnet.TRouters
}

func (pr *PingRouter) PreHandle(request tface.TRequest) {
	fmt.Println("Call PingRouter--PreHandle...")
	_, err := request.GetConn().GetTCPConnection().Write([]byte("Before Ping Router...\n"))
	if err != nil {
		fmt.Println("Call Before PingRouter Error!")
	}
}

func (pr *PingRouter) Handler(request tface.TRequest) {
	fmt.Println("Call PingRouter--Handler...")
	_, err := request.GetConn().GetTCPConnection().Write([]byte("Ping Ping Ping...\n"))
	if err != nil {
		fmt.Println("Call PingRouter Error!")
	}
}

func (pr *PingRouter) PostHandle(request tface.TRequest) {
	fmt.Println("Call PingRouter--PostHandle...")
	_, err := request.GetConn().GetTCPConnection().Write([]byte("After Ping Router...\n"))
	if err != nil {
		fmt.Println("Call After PingRouter Error!\n")
	}

}

func main() {
	s := tnet.NewServer("demoV0.3")
	s.AddRouter(&PingRouter{})
	s.ServerRun()
}
