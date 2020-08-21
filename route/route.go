package route

import (
	blogApi "SkyPalace/app/blog/api"
	ttApi "SkyPalace/app/tianting/api"
	ttModelRoute "SkyPalace/app/tianting/route"
	"github.com/gin-gonic/gin"
)

func RouteCtrl(r *gin.Engine) {
	BlogRoute := r.Group("/blog")
	{
		BlogRoute.POST("user/login", blogApi.Login)
		BlogRoute.POST("user/register", blogApi.UserRegister)

	}

	TTroute := r.Group("/tt")
	{
		//TTroute.Use(middleware.AuthRequired())

		// 南天门
		TTroute.POST("user/login", ttApi.Login)

		// 封神榜
		TTroute.POST("user/register", ttApi.UserRegister)

		// 瑶池
		ttModelRoute.Yaochi(TTroute)

		//// 凌霄殿
		//Lingxiao(TTroute)
		//
		//// 兜率宫
		//DouLv(TTroute)
	}
}
