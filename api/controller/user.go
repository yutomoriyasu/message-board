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
	return User{
		userInteractor: userInteractor,
	}
}

func (u *User) CreateUser(ctx echo.Context) error {
	var param openapi.CreateUserParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	name := user.NewName(param.Name)

	email, err := user.NewEmail(param.Email)
	if err != nil {
		return err
	}

	user, err := u.userInteractor.CreateUser(ctx.Request().Context(), name, email)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, openapi.NewUser(user))
}
