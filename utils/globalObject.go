package utils

import (
	"TCPDemo/tface"
	"encoding/json"
	"io/ioutil"
)

// GlobalObj 定义一个存储一切有关框架的全局参数的对象，提供给其他模块使用
// 大部分参数通过 json 文件由用户进行配置
type GlobalObj struct {
	//server
	TCPServer tface.TServer
	Host      string
	TCPPort   int
	Name      string // 服务器名称

	Version    string
	MaxConn    int    // 服务器主机允许的最大连接数
	MaxPackage uint32 // 当前数据包的最大值
}

// GlobalObject 定义一个全局的 object 对象
var GlobalObject *GlobalObj

// 提供一个 init 方法，初始化全局对象

// Load 从配置文件中加载的相关参数
func (g *GlobalObj) Load() {
	data, err := ioutil.ReadFile("conf/TCPDemo.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
func init() {
	// 如果没有加载配置文件，就使用下面的默认值 2023-09-06
	GlobalObject = &GlobalObj{
		Name:       "TCPDemo",
		Host:       "0.0.0.0",
		TCPPort:    8090,
		Version:    "0.1",
		MaxConn:    1000,
		MaxPackage: 4096,
	}

	// 尝试从配置文件中读取相关信息
	GlobalObject.Load()
}
