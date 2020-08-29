package database

import (
	"MySystem/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MYSQLDBTT 数据库链接单例
var MYSQLDBTT *gorm.DB

// Database 在中间件中初始化mysql链接
func MysqlInit(connString string, DBtype string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}

	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	if DBtype == "MYSQL_TT" {
		MYSQLDBTT = db
	}

}

func GetTTDB() *gorm.DB {
	return MYSQLDBTT
}
