package route

//import (
//	"MySystem/api"
//	"MySystem/middleware"
//	"github.com/gin-gonic/gin"
//)
//
//func Lingxiao(r *gin.RouterGroup) {
//	lx := r.Group("/lingxiao")
//	{
//		lx.GET("/", api.Lingxiao) // 凌霄主页
//		lx.GET("PM25", api.Air)
//
//		lx_auth := lx.Group("")
//		lx_auth.Use(middleware.AuthRequired())
//		{
//			lx_auth.GET("meeting", api.Meeting) // 开会
//			lx_auth.GET("event", api.Event)
//		}
//
//	}
//
//}
