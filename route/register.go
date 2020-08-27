package route

import (
	"github.com/gin-gonic/gin"
)

func NewRegister() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	//r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	//r.Use(middleware.Cors())
	//r.Use(middleware.CurrentUser())

	// 首页
	r.GET("", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	RouteCtrl(r)

	return r
}
