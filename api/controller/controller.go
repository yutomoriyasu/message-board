package controller

import "message-board/usecase"

type Controller struct {
	User
}

func NewController(userInteractor usecase.IUserUsecase) *Controller {
	return &Controller{
		User: NewUser(userInteractor),
	}
}
