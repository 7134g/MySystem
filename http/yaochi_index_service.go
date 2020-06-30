package http

import (
	"SkyPalace/database"
	"SkyPalace/model"
	"SkyPalace/serializer"
	"github.com/gin-gonic/gin"
)

// 基本数据
type YcBaseService struct {
}

// 树林数据
type YcPeachStatusService struct {
}

func (self YcBaseService) BaseInfo(c *gin.Context) serializer.Response {
	var v model.YcBase
	err := database.GetDB().First(&v)
}
