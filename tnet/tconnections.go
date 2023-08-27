package tnet

import (
	"TCPDemo/tface"
	"fmt"
	"net"
)

// 链接模块

type Tconnections struct {
	//当前连接的TCP套接字
	Conn *net.TCPConn
	//当前连接的ID
	ConnID uint32
	//当前连接状态
	isClosed bool
	//当前连接绑定的业务方法
	handleAPI tface.HandleFunc
	//告知当前连接已经退出/停止的channel
	Exitchan chan bool
	//该链接处理的 Router
	Router tface.TRouter
}

// 初始化连接的方法
func NewConnection(conn *net.TCPConn, id uint32, router tface.TRouter) *Tconnections {
	c := &Tconnections{
		Conn:     conn,
		ConnID:   id,
		Router:   router,
		isClosed: false,
		Exitchan: make(chan bool, 1),
	}
	return c
}

// 连接的读业务
func (c *Tconnections) StartReader() {
	fmt.Println("[INFO] Reader Goroutine is Running...")
	defer fmt.Println("[INFO] conID = ", c.ConnID, "Is Stopping", " Reader Will Stop Later,Remote Addr is ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端数据到BUFF,目前最大512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("[Error] Reader Can Not Read Data From Buffer:", err)
			continue
		}
		// 得到一个 Request 对象
		req := TRequest{
			conn: c,
			data: buf,
		}

		// 调用当前链接所绑定的 Router 2023-08-27
		go func(request tface.TRequest) {
			c.Router.PreHandle(request)
			c.Router.Handler(request)
			c.Router.PostHandle(request)

		}(&req)
		//调用当前连接所绑定的HandleAPI
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("[Error] HandleAPI Run Error :", err)
		//	break
		//}
	}
}

func (c *Tconnections) Start() {
	fmt.Println("[INFO] Connection Is Start...Connection ID = ", c.ConnID)
	//启动从当前连接读数据的业务
	go c.StartReader()
	//TODO: 启动从当前连接写数据的业务
}
func (c *Tconnections) Stop() {
	fmt.Println("[INFO] Connection is Stop...Connection ID = ", c.ConnID)

	//如果当前连接已经关闭
	if c.isClosed == true {
		return
	}
	//关闭链接
	c.isClosed = true
	c.Conn.Close()
	//回收资源
	close(c.Exitchan)
}
func (c *Tconnections) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Tconnections) GetConnectionID() uint32 {
	return c.ConnID
}
func (c *Tconnections) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
func (c *Tconnections) Send(data []byte) error {
	return nil
}
