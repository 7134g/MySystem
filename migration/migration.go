package migration

import (
	ttModel "SkyPalace/app/tianting/model"
	"SkyPalace/database"
)

//执行数据迁移

func Migration() {
	// 自动迁移模式
	db := database.GetDB()
	db.AutoMigrate(&ttModel.TTUser{})
	db.AutoMigrate(&ttModel.YcBase{})
	db.AutoMigrate(&ttModel.YcPeachStatus{})
}
