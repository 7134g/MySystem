package serializer

import "MySystem/model"

// User 用户序列化器
type User struct {
	ID          uint   `json:"id"`
	Username    string `json:"user_name"`
	Nickname    string `json:"nickname"`
	Status      string `json:"status"`
	Avatar      string `json:"avatar"`
	CreatedAt   int64  `json:"created_at"`
	Permissions int8   `json:"permissions"`
}

func BuildUser(user model.User) User {
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

func BuildUserResponse(user model.User) Response {
	return Response{Data: BuildUser(user)}
}
