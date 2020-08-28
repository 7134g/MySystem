package api

import (
	"MySystem/app/blog/service"
	"MySystem/lib"
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	s := service.ArticleService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.CreateArticle(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}

func DelArticle(c *gin.Context) {
	s := service.ArticleService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.DelArticle(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}

func UpdateArticle(c *gin.Context) {
	s := service.ArticleService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.UpdateArticle(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, lib.ErrorResponse(err))
	}
}
