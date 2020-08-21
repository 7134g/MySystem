package api

import (
	"SkyPalace/app/tianting/service"
	"SkyPalace/lib"
	"github.com/gin-gonic/gin"
)

func Yaochi(c *gin.Context) {
	s := service.YcBaseService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.BaseInfo(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}

//func PeachForest(c *gin.Context) {
//	s := http.Service{}
//	if err := c.ShouldBind(&s); err == nil {
//		res := s.FuncMethod(c)
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
//
//func PeriAction(c *gin.Context) {
//	s := http.Service{}
//	if err := c.ShouldBind(&s); err == nil {
//		res := s.FuncMethod(c)
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
