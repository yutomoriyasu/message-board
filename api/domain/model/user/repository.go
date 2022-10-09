package user

import "context"

type IRepository interface {
	Create(context.Context, *User) (*User, error)
	Find(context.Context) (Users, error)
}
