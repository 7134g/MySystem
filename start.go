package main

import (
	"SkyPalace/config"
	"SkyPalace/server"
	"SkyPalace/util"
)

func main() {
	// 从配置文件读取配置
	config.Init()

	// 装载路由
	s := server.NewRegister()
	err := s.Run(":3000")
	if err != nil {
		util.Log().Panic("服务器启动失败", err)
	}
}
