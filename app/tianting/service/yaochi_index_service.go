package service

import (
	"MySystem/app/tianting/model"
	"MySystem/database"
	"MySystem/lib"
	"MySystem/util"
	"github.com/gin-gonic/gin"
)

// 基本数据
type YcBaseService struct {
}

// 树林数据
type YcPeachStatusService struct {
}

func (self YcBaseService) BaseInfo(c *gin.Context) lib.Response {
	var v model.YcBase
	err := database.GetTTDB().Exec("select sum(local) as local,").Find(&v).Error
	if err != nil {
		util.Log().Error(err.Error())
	}
	return lib.Response{
		Code: 1111111,
		Msg:  "test",
		Data: v,
	}

}
