package model

import (
	"MySystem/database"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户模型
type BGUser struct {
	ID             uint
	Username       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string
	CreatedAt      time.Time
	Permissions    int8
}

const PassWordCost = 12

// GetUser 用ID获取用户
func GetUser(ID interface{}) (BGUser, error) {
	var user BGUser
	result := database.GetTTDB().First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *BGUser) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *BGUser) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
