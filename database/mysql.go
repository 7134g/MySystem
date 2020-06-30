package database

import (
	"SkyPalace/util"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MYSQLDB 数据库链接单例
var MYSQLDB *gorm.DB

// Database 在中间件中初始化mysql链接
func MysqlInit(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	MYSQLDB = db

}

func GetDB() *gorm.DB {
	return MYSQLDB
}
