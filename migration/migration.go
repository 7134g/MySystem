package migration

import (
	ttModel "MySystem/app/tianting/model"
	"MySystem/database"
)

//执行数据迁移

func Migration() {
	// 自动迁移模式
	db := database.GetTTDB()
	db.AutoMigrate(&ttModel.TTUser{})
	db.AutoMigrate(&ttModel.YcBase{})
	db.AutoMigrate(&ttModel.YcPeachStatus{})
}
