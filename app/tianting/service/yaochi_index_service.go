package service

import (
	"MySystem/app/tianting/model"
	"MySystem/database"
	"MySystem/serializer"
	"MySystem/util"
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
	err := database.GetDB().Exec("select sum(local) as local,").Find(&v).Error
	if err != nil {
		util.Log().Error(err.Error())
	}
	return serializer.Response{
		Code: 1111111,
		Msg:  "test",
		Data: v,
	}

}
