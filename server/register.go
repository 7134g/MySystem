package server

import (
	"SkyPalace/middleware"
	"SkyPalace/route"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRegister() *gin.Engine {
	r := gin.Default()
	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	r.GET("", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	route.RouteCtrl(r)

	return r
}
