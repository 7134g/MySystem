package route

import (
	"SkyPalace/app/tianting/api"
	"github.com/gin-gonic/gin"
)

func Yaochi(r *gin.RouterGroup) {
	yc := r.Group("/yaochi")
	{
		yc.GET("/", api.Yaochi) // 瑶池主页
		//yc.GET("peachforest", api.PeachForest) // 瑶池蟠桃林
		//yc.GET("periaction", api.PeriAction)   // 瑶池宫
	}

}
