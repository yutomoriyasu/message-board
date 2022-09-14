package openapi

import (
	"message-board/domain/model/user"
)

func NewUser(user *user.User) User {
	return User{
		Id:    user.ID.Uint64(),
		Name:  user.Name.String(),
		Email: user.Email.String(),
	}
}
