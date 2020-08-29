package service

import (
	"MySystem/app/blog/model"
	"MySystem/app/blog/serializer"
	"MySystem/database"
	"MySystem/lib"
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (s *UserLoginService) setSession(c *gin.Context, user model.BGUser) {
	session := sessions.Default(c)
	session.Clear()
	session.Set("user_id", user.ID)
	_ = session.Save()
}

// Login 用户登录函数
func (s *UserLoginService) Login(c *gin.Context) lib.Response {
	var user model.BGUser
	var err error

	userCol := database.MONGODBBG.Collection("user")
	ctx := context.TODO()

	err = userCol.FindOne(ctx, bson.D{{"username", s.UserName}}).Decode(&user)
	if err != nil {
		return lib.ParamErr("账号或密码错误", err)
	}

	//if err := database.MONGODBBG.Where("username = ?", s.UserName).First(&user).Error; err != nil {
	//	return lib.ParamErr("账号或密码错误", err)
	//}

	if user.CheckPassword(s.Password) == false {
		return lib.ParamErr("账号或密码错误", nil)
	}

	s.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
