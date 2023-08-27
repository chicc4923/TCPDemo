package tnet

import (
	"TCPDemo/tface"
	"fmt"
	"net"
)

type TServers struct {
	//服务器名称
	Name string
	//服务器监听IP
	IP string
	//服务器监听端口
	Port int
	//服务器绑定IP版本
	IPVersion string
	//给当前的 Server 添加一个 Router, Server 注册的连接对应的处理业务
	Router tface.TRouter
}

//实例化接口层的TServer

func (s *TServers) Start() {
	fmt.Printf("[START] Server Listening at %s:%d", s.IP, s.Port)

	go func() {

		//1. 获取TCP的ADDR
		Addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("[Error] Server Can Not Resolve TCP Address:", err)
		}

		//2. 监听服务地址
		Listener, err := net.ListenTCP(s.IPVersion, Addr)
		if err != nil {
			fmt.Println("[Error] Server Can Not Listen TCP Address:", err)
			return
		}

		fmt.Println("[INFO] Start TCPDemo Success,", s.Name, "Is Listening")
		var cid uint32 = 0

		//3. 阻塞的等待新的客户端连接进入,然后处理相关逻辑
		for {
			// 如果有客户端连接，阻塞会返回
			conn, err := Listener.AcceptTCP()
			if err != nil {
				fmt.Println("[Waring] Accept Client Connection Error:", err)
				continue
			}
			//将处理新链接的业务方法和业务进行绑定，得到连接模块。
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			//启动连接模块
			go dealConn.Start()

			//客户端连接后，开始进行对应的逻辑处理
			//基本回写业务（最大 521 字节）

			/*	go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("[ERROR] Server Can Not Get Msg From Buffer", err)
						continue
					}
					//回写功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("[ERROR] Server Write Back Error:", err)
					}
				}
			}() */

		}
	}()

}

func (s *TServers) Stop() {}
func (s *TServers) ServerRun() {
	s.Start()

	//做一些启动服务器后的一些其他工作

	//阻塞服务
	select {}
}

// 初始化Server
func NewServer(name string) *TServers {
	s := &TServers{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}

// 定义当前客户端连接所绑定的业务逻辑(Handle API,以后需要用户自定义)
//func CallBack(conn *net.TCPConn, data []byte, cnt int) error {
//	// 回显业务
//	fmt.Println("[INFO] Connection Handle Back To Clinet...")
//	if _, err := conn.Write(data[:cnt]); err != nil {
//		fmt.Println("[ERROR] Write Back Buffer Error:", err)
//		return err
//	}
//	return nil
//}

// 添加一个路由方法 2023-08-27
func (s *TServers) AddRouter(router tface.TRouter) {
	s.Router = router
	fmt.Println("Add Router Success!")
}
