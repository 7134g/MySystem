package config

import (
	"SkyPalace/database"
	"SkyPalace/model"
	"SkyPalace/util"
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		util.Log().Panic("环境文件错误", err)
	}

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	database.MysqlInit(os.Getenv("MYSQL_DSN"))
	model.Migration()
	database.RedisInit()

}
