package service

import (
	"MySystem/app/blog/model"
	"MySystem/app/blog/serializer"
	"MySystem/database"
	"MySystem/lib"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
	Permissions     int8
}

func (self *UserRegisterService) Register() lib.Response {
	user := model.BGUser{
		Nickname:    self.Nickname,
		Username:    self.UserName,
		Permissions: 1,
	}

	if err := self.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(self.Password); err != nil {
		return lib.Err(
			lib.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := database.GetBGDB().Create(&user).Error; err != nil {
		return lib.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}

func (service *UserRegisterService) valid() *lib.Response {
	if service.PasswordConfirm != service.Password {
		return &lib.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	database.GetBGDB().Model(&model.BGUser{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &lib.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}

	count = 0
	database.GetBGDB().Model(&model.BGUser{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &lib.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}
