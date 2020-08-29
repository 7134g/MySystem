package database

import (
	"os"
	"strings"
)

func DataBaseInit() {

	// 数据库初始化
	// 获取Mongo链接参数
	ClientURI := os.Getenv("MONGO_CLIENT")
	DataNames := strings.Fields(os.Getenv("MONGO_DBNAME"))
	// 连接Mongo数据库
	MongoInit(ClientURI, DataNames)

	// 链接redis数据库
	RedisInit()
}
