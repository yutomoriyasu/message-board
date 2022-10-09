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

func NewUsers(us user.Users) Users {
	users := make([]User, len(us))
	for i, u := range us {
		users[i] = User{
			Id:    u.ID.Uint64(),
			Name:  u.Name.String(),
			Email: u.Email.String(),
		}
	}
	return Users{
		Users: users,
	}
}
