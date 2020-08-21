package api

import (
	"SkyPalace/lib"
	"SkyPalace/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	s := service.UserLoginService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}

func UserRegister(c *gin.Context) {
	s := service.UserRegisterService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}
