package model

// 基本数据
type YcBase struct {
	local      string // 十六进制
	PeachTree  int64  // 桃树
	Peach      int64  // 桃子
	area       int64  // 面积
	ReikiLevel int64  // 灵气浓度
	YcPeachStatus
}

// 树林数据
type YcPeachStatus struct {
	ImmaturePeach     int64 // 未成熟桃子
	ImmaturePeachTree int64 // 未成熟桃树
	YcPeri            int64 // 瑶池仙女数
	TodayHarvest      int64 // 今日收获量
}
