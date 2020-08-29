package service

import (
	"MySystem/app/blog/model"
	"MySystem/app/blog/serializer"
	"MySystem/database"
	"MySystem/lib"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
	Permissions     int8
}

func (s *UserRegisterService) Register() lib.Response {
	user := model.BGUser{
		Username:    s.UserName,
		Nickname:    s.Nickname,
		Permissions: 1,
	}
	userCol := database.MONGODBBG.Collection("user")

	if err := s.valid(userCol); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(s.Password); err != nil {
		return lib.Err(
			lib.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if _, err := userCol.InsertOne(context.TODO(), user); err != nil {
		return lib.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}

func (s *UserRegisterService) valid(col *mongo.Collection) *lib.Response {
	var err error
	var user model.BGUser

	if s.PasswordConfirm != s.Password {
		return &lib.Response{
			Code: lib.CodeParamErr,
			Msg:  "两次输入的密码不相同",
		}
	}

	// 检查用户是否存在
	filter := bson.D{
		{"username", s.UserName},
	}
	err = col.FindOne(context.Background(), filter).Decode(&user)
	if err == nil {
		return &lib.Response{
			Code: lib.CodeParamErr,
			Msg:  "用户已存在",
		}
	}

	//count := 0
	//database.MONGODBBG.Model(&model.BGUser{}).Where("nickname = ?", s.Nickname).Count(&count)
	//if count > 0 {
	//	return &lib.Response{
	//		Code: lib.CodeParamErr,
	//		Msg:  "昵称被占用",
	//	}
	//}
	//

	return nil
}
