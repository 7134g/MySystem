package route

import (
	blogApi "MySystem/app/blog/api"
	"github.com/gin-gonic/gin"
)

func Blogroute(r *gin.Engine) {
	bg := r.Group("/blog")
	{
		bg.POST("user/login", blogApi.Login)
		bg.POST("user/register", blogApi.UserRegister)

	}
}
