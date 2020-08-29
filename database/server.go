package database

import (
	"os"
	"strings"
)

func DBInit() {
	// 数据库初始化

	// 获取Mongo链接参数
	ClientURI := os.Getenv("MONGO_CLIENT")
	DataNames := strings.Fields(os.Getenv("MONGO_DBNAME"))
	// 连接Mongo数据库
	MongoInit(ClientURI, DataNames)

	// 获取Mysql链接参数
	MYSQlURI := os.Getenv("MYSQL_TT")
	// 链接Mysql数据库
	MysqlInit(MYSQlURI, "MYSQL_TT")

	// 链接redis数据库
	RedisInit()
}
