package model

import (
	"SkyPalace/database"
)

//执行数据迁移

func Migration() {
	// 自动迁移模式
	db := database.GetDB()
	db.AutoMigrate(&User{})
}
