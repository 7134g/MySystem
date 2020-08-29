package route

import (
	BGR "MySystem/app/blog/route"
	TTR "MySystem/app/tianting/route"
	"github.com/gin-gonic/gin"
)

func Control(r *gin.Engine) {
	BGR.Blogroute(r)
	TTR.TTroute(r)
	//DFR.DFroute(r)
}
