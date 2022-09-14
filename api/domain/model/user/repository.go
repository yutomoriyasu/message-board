package user

import "context"

type IRepository interface {
	Create(context.Context, *User) (ID, error)
}
