package usecase

import (
	"context"
	"message-board/domain/model/user"
)

type IUserUsecase interface {
	CreateUser(context.Context, user.Name, user.Email) (*user.ID, error)
}

type userUsecase struct {
	userRepo user.IRepository
}

func NewUser(u user.IRepository) IUserUsecase {
	return &userUsecase{userRepo: u}
}

func (u *userUsecase) CreateUser(ctx context.Context, name user.Name, email user.Email) (*user.ID, error) {
	user := user.NewUserForCreate(name, email)

	userID, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &userID, nil
}
