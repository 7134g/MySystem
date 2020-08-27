package main

import (
	"MySystem/config"
	"MySystem/server"
	"MySystem/util"
)

func main() {
	// 从配置文件读取配置
	config.Init()

	// 装载路由
	s := server.NewRegister()
	err := s.Run(":8500")
	if err != nil {
		util.Log().Panic("服务器启动失败", err)
	}
}
