package api

import (
	"SkyPalace/http"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	s := http.UserLoginService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserRegister(c *gin.Context) {
	s := http.UserRegisterService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
