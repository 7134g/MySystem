package api

import (
	"MySystem/app/blog/service"
	"MySystem/lib"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	s := service.NarBarService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Home(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}

func ShowArticle(c *gin.Context) {
	s := service.NarBarService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.ShowArticle(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}

func Classfication(c *gin.Context) {
	s := service.NarBarService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Classfication(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}

func TimeLine(c *gin.Context) {
	s := service.NarBarService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.TimeLine(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}

func JianLi(c *gin.Context) {
	s := service.NarBarService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.JianLi(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}
