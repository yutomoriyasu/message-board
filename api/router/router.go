package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	e *echo.Echo
}

func NewRouter(e *echo.Echo) *Router {
	return &Router{e: e}
}

func (r Router) Init(port string) {
	r.e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, World! <3")
	})

	r.e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
	r.e.Debug = true
	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())

	r.e.Logger.Fatal(r.e.Start(":" + port))
}
