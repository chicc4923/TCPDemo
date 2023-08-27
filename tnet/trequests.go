package tnet

import "TCPDemo/tface"

type TRequest struct {
	//已经和客户端建立好的连接
	conn tface.TConnection
	//客户端请求的数据
	data []byte
}

// 得到当前连接
func (r *TRequest) GetConn() tface.TConnection {
	return r.conn
}

// 得到请求的消息数据
func (r *TRequest) GetData() []byte {
	return r.data
}
