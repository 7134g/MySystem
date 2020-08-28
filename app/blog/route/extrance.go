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

		bg.GET("home", blogApi.Home)
		bg.GET("ShowArticles", blogApi.ShowArticle)
		bg.GET("classfication", blogApi.Classfication)
		bg.GET("TimeLine", blogApi.TimeLine)
		bg.GET("JianLi", blogApi.JianLi)

		// 需要登录保护的
		authed := bg.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// 创建文章
			bg.POST("NewArticle", blogApi.CreateArticle)
			// 删除文章
			bg.DELETE("DelArticle", blogApi.DelArticle)
			// 更新文章
			bg.POST("ChangeArticle", blogApi.UpdateArticle)
		}

	}

}
