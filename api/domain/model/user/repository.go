package user

import "context"

type IRepository interface {
	Create(context.Context, *User) (*User, error)
	Find(context.Context) (Users, error)
	GetByID(context.Context, ID) (*User, error)
	Update(context.Context, *User) error
	Delete(context.Context, ID) error
}
