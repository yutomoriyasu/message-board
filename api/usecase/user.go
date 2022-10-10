package usecase

import (
	"context"
	"message-board/domain/model/user"
)

type IUserUsecase interface {
	CreateUser(context.Context, user.Name, user.Email) (*user.User, error)
	Find(context.Context) (user.Users, error)
	GetUserByID(context.Context, user.ID) (*user.User, error)
	UpdateUser(context.Context, *user.User) error
	DeleteUser(context.Context, user.ID) error
}

type userUsecase struct {
	userRepo user.IRepository
}

func NewUser(userRepo user.IRepository) IUserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) CreateUser(ctx context.Context, name user.Name, email user.Email) (*user.User, error) {
	userForCreate := user.NewUserForCreate(name, email)

	user, err := u.userRepo.Create(ctx, userForCreate)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Find(ctx context.Context) (user.Users, error) {
	users, err := u.userRepo.Find(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) GetUserByID(ctx context.Context, id user.ID) (*user.User, error) {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, user *user.User) error {
	err := u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, id user.ID) error {
	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
