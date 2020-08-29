package serializer

import (
	"MySystem/app/blog/model"
	"MySystem/lib"
)

// User 用户序列化器
type User struct {
	ID          uint   `bson:"id"`
	Username    string `bson:"user_name"`
	Nickname    string `bson:"nickname"`
	Status      string `bson:"status"`
	Avatar      string `bson:"avatar"`
	CreatedAt   int64  `bson:"created_at"`
	Permissions int8   `bson:"permissions"`
}

func BuildUser(user model.BGUser) User {
	return User{
		ID:          user.ID,
		Username:    user.Username,
		Nickname:    user.Nickname,
		Status:      user.Status,
		Avatar:      user.Avatar,
		CreatedAt:   user.CreatedAt.Unix(),
		Permissions: user.Permissions,
	}
}

func BuildUserResponse(user model.BGUser) lib.Response {
	return lib.Response{Data: BuildUser(user)}
}
