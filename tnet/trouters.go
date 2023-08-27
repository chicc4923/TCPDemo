package tnet

import "TCPDemo/tface"

// 定义该Router实现时先嵌入这个router基类，然后根据需要重写这个基类的方法
type TRouters struct{}

// 这里之所以所有的方法都为空，是因为有的 router 不需要Pre 和 POST 两个方法
// 所以 Router 全部继承 TRouters 时不需要实现 Pre 和 Post 两个方法
// 处理业务之前的方法
func (tr *TRouters) PreHandle(request tface.TRequest) {}

// 处理业务的方法
func (tr *TRouters) Handler(request tface.TRequest) {}

// 处理业务之后的方法
func (tr *TRouters) PostHandle(request tface.TRequest) {}
