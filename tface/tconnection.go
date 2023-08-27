package tface

import "net"

type TConnection interface {
	//启动连接,
	Start()
	//停止连接
	Stop()
	//获取当前连接绑定的TCP SOCKET connection
	GetTCPConnection() *net.TCPConn
	//获取当前连接的ID
	GetConnectionID() uint32
	//获取远程客户端连接的地址和端口
	GetRemoteAddr() net.Addr
	//发送数据
	Send(data []byte) error
}

// 定义一个处理
type HandleFunc func(*net.TCPConn, []byte, int) error
