package route

import (
	"SkyPalace/api"
	"github.com/gin-gonic/gin"
)

func TianTingCtrl(r *gin.Engine) {
	TTroute := r.Group("/tt")
	//TTroute.Use(middleware.AuthRequired())

	// 南天门
	TTroute.POST("user/login", api.Login)

	// 封神榜
	TTroute.POST("user/register", api.UserRegister)

	// 瑶池
	Yaochi(TTroute)

	//// 凌霄殿
	//Lingxiao(TTroute)
	//
	//// 兜率宫
	//DouLv(TTroute)

}
