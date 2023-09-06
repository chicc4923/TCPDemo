package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	//1.与服务端建立连接，得到一个conn
	conn, err := net.Dial("tcp", "127.0.0.1:8989")
	if err != nil {
		fmt.Printf("connection failed:", err)
	}

	for {
		// 连接调用write，写数据
		_, err := conn.Write([]byte("hello,zinx V0.1\n"))
		if err != nil {
			fmt.Println("write conn failed:", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error:", err)
			return
		}
		fmt.Printf("Server call back:%s,cnt:%d\n", buf, cnt)

		//cpu阻塞
		time.Sleep(1 * time.Second)
	}
}
