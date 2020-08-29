package route

import (
	"github.com/gin-gonic/gin"
)

func NewRegister() *gin.Engine {
	r := gin.Default()

	// 首页
	r.GET("", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	// 路由初始化
	Control(r)

	return r
}
