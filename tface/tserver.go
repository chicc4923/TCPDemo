package tface

type TServer interface {
	//服务器启动
	Start()
	//服务器停止
	Stop()
	//服务器运行
	ServerRun()
}
