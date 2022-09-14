//go:build wireinject
// +build wireinject

package main

import (
	"message-board/controller"
	"message-board/infrastructure/postgres"
	"message-board/router"
	"message-board/usecase"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var Set = wire.NewSet(
	postgres.Set,
	usecase.Set,
	controller.NewController,
	router.NewRouter,
)

func InitRouter(e *echo.Echo, db postgres.DB) *router.Router {
	wire.Build(Set)
	return &router.Router{}
}
