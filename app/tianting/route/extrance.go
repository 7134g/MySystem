package route

import (
	ttApi "MySystem/app/tianting/api"
	"MySystem/app/tianting/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func TTroute(r *gin.Engine) {
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	tt := r.Group("/tt")
	{
		//TTroute.Use(middleware.AuthRequired())

		// 南天门
		tt.POST("user/login", ttApi.Login)

		// 封神榜
		tt.POST("user/register", ttApi.UserRegister)

		// 瑶池
		Yaochi(tt)

		//// 凌霄殿
		//Lingxiao(tt)
		//
		//// 兜率宫
		//DouLv(tt)

		// 需要登录保护的
		authed := tt.Group("/")
		authed.Use(middleware.AuthRequired())
		{

		}

	}
}
