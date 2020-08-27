package route

import (
	blogApi "MySystem/app/blog/api"
	"MySystem/app/blog/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func Blogroute(r *gin.Engine) {
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	bg := r.Group("/blog")
	{
		bg.POST("user/login", blogApi.Login)
		bg.POST("user/register", blogApi.UserRegister)

		// 需要登录保护的
		authed := bg.Group("/")
		authed.Use(middleware.AuthRequired())
		{

		}

	}

}
