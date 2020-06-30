package http

import (
	"SkyPalace/database"
	"SkyPalace/model"
	"SkyPalace/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
	Permissions     int8
}

func (self *UserRegisterService) Register() serializer.Response {
	user := model.User{
		Nickname:    self.Nickname,
		Username:    self.UserName,
		Status:      model.Active,
		Permissions: 1,
	}

	if err := self.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(self.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := database.GetDB().Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}

func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	database.GetDB().Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}

	count = 0
	database.GetDB().Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}
