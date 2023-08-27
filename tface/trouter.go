package tface

type TRouter interface {
	//路由抽象接口
	//路由里的数据均为TRequests

	//处理业务之前的方法
	PreHandle(request TRequest)
	//处理业务的方法
	Handler(request TRequest)
	//处理业务之后的方法
	PostHandle(request TRequest)
}
