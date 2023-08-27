package tface

// 把客户端请求的链接信息和数据包装到一个 Request 中
type TRequest interface {
	//得到当前连接
	GetConn() TConnection
	//得到请求的消息数据
	GetData() []byte
}
