package route

import (
	ttApi "MySystem/app/tianting/api"
	"github.com/gin-gonic/gin"
)

func TTroute(r *gin.Engine) {
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
	}
}
