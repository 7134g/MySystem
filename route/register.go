package route

import (
	"MySystem/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewRegister() *gin.Engine {
	r := gin.Default()

	// 从本地读取环境变量
	godotenv.Load()

	//// 读取翻译文件
	//if err := config.LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
	//	panic(err)
	//}

	// 首页
	r.GET("", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	// 路由初始化
	RouteCtrl(r)

	// 数据库初始化
	database.DataBaseInit()

	return r
}
