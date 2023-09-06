package tface

type TServer interface {
	// Start 服务器启动
	Start()
	// Stop 服务器停止
	Stop()
	// ServerRun 服务器运行
	ServerRun()

	// AddRouter 路由功能：给当前的服务注册一个路由方法,供客户端的连接处理使用
	AddRouter()
}
