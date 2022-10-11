package router

import (
	"message-board/controller"
	"message-board/controller/openapi"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	e *echo.Echo
	c *controller.Controller
}

func NewRouter(e *echo.Echo, c *controller.Controller) *Router {
	openapi.RegisterHandlers(e.Group("/api"), c)
	return &Router{e, c}
}

func (r Router) Init(port string) {
	r.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! <3")
	})
	r.e.Debug = true
	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())

	r.e.Logger.Fatal(r.e.Start(":" + port))
}
