package controller

import (
	"message-board/controller/openapi"
	"message-board/domain/model/user"
	"message-board/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	userInteractor usecase.IUserUsecase
}

func NewUser(userInteractor usecase.IUserUsecase) User {
	return User{userInteractor}
}

func (u *User) CreateUser(ctx echo.Context) error {
	var param openapi.CreateUserParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	name := user.NewName(string(param.Name))

	email, err := user.NewEmail(string(param.Email))
	if err != nil {
		return err
	}

	user, err := u.userInteractor.CreateUser(
		ctx.Request().Context(), name, email,
	)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, openapi.NewUser(user))
}

func (u *User) GetUsers(ctx echo.Context) error {
	users, err := u.userInteractor.Find(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, openapi.NewUsers(users))
}

func (u *User) GetUser(ctx echo.Context, userId openapi.UserId) error {
	user, err := u.userInteractor.GetUserByID(
		ctx.Request().Context(), user.NewID(uint64(userId)),
	)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, openapi.NewUser(user))
}

func (u *User) UpdateUser(ctx echo.Context, userId openapi.UserId) error {
	var param openapi.UpdateUserParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}
	email, err := user.NewEmail(param.Email)
	if err != nil {
		return err
	}
	user := user.NewUser(
		user.NewID(uint64(userId)),
		user.NewName(param.Name),
		email,
	)
	err = u.userInteractor.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (u *User) DeleteUser(ctx echo.Context, userId openapi.UserId) error {
	err := u.userInteractor.DeleteUser(
		ctx.Request().Context(),
		user.NewID(uint64(userId)),
	)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}
