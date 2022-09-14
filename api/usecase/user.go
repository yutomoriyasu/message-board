package usecase

import (
	"context"
	"message-board/domain/model/user"
)

type IUserUsecase interface {
	CreateUser(context.Context, user.Name, user.Email) (*user.User, error)
}

type userUsecase struct {
	userRepo user.IRepository
}

func NewUser(u user.IRepository) IUserUsecase {
	return &userUsecase{userRepo: u}
}

func (u *userUsecase) CreateUser(ctx context.Context, name user.Name, email user.Email) (*user.User, error) {
	user := user.NewUserForCreate(name, email)

	user, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
