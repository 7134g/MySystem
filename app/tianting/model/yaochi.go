package model

import "github.com/jinzhu/gorm"

// 基本数据
type YcBase struct {
	gorm.Model
	Local      string // 十六进制
	PeachTree  int64  // 桃树
	Peach      int64  // 桃子
	area       int64  // 面积
	ReikiLevel int64  // 灵气浓度
}

// 树林数据
type YcPeachStatus struct {
	gorm.Model
	ImmaturePeach     int64 // 未成熟桃子
	ImmaturePeachTree int64 // 未成熟桃树
	YcPeri            int64 // 瑶池仙女数
	TodayHarvest      int64 // 今日收获量
}
