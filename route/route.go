package route

import (
	BGroute "MySystem/app/blog/route"
	TTroute "MySystem/app/tianting/route"
	"github.com/gin-gonic/gin"
)

func RouteCtrl(r *gin.Engine) {
	BGroute.Blogroute(r)
	TTroute.TTroute(r)

}
