package openapi

import (
	"message-board/domain/model/user"
)

func NewUser(user *user.User) User {
	return User{
		Id: Id(user.ID),
		UserProps: UserProps{
			Name:  user.Name.String(),
			Email: user.Email.String(),
		},
	}
}

func NewUsers(us user.Users) Users {
	users := make([]User, len(us))
	for i, u := range us {
		users[i] = NewUser(u)
	}
	return Users{
		Users: users,
	}
}
