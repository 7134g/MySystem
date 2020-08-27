package database

import (
	"MySystem/util"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MYSQLDBTT 数据库链接单例
var MYSQLDBTT *gorm.DB
var MYSQLDBBG *gorm.DB

// Database 在中间件中初始化mysql链接
func MysqlInit(connString string, DBtype string) {
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

	if DBtype == "MYSQL_TT" {
		MYSQLDBTT = db
	}

	if DBtype == "MYSQL_BG" {
		MYSQLDBBG = db
	}
}

func GetTTDB() *gorm.DB {
	return MYSQLDBTT
}

func GetBGDB() *gorm.DB {
	return MYSQLDBBG
}
