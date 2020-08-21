package service

import (
	"SkyPalace/database"
	"SkyPalace/model"
	"SkyPalace/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (self *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	_ = s.Save()
}

// Login 用户登录函数
func (self *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := database.MYSQLDB.Where("username = ?", self.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", err)
	}

	if user.CheckPassword(self.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	self.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
