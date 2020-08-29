package service

import (
	"MySystem/app/tianting/model"
	"MySystem/app/tianting/serializer"
	"MySystem/database"
	"MySystem/lib"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (s *UserLoginService) setSession(c *gin.Context, user model.TTUser) {
	session := sessions.Default(c)
	session.Clear()
	session.Set("user_id", user.ID)
	_ = session.Save()
}

// Login 用户登录函数
func (s *UserLoginService) Login(c *gin.Context) lib.Response {
	var user model.TTUser

	if err := database.MYSQLDBTT.Where("username = ?", s.UserName).First(&user).Error; err != nil {
		return lib.ParamErr("账号或密码错误", err)
	}

	if user.CheckPassword(s.Password) == false {
		return lib.ParamErr("账号或密码错误", nil)
	}

	s.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
