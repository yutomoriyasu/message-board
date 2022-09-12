package main

import (
	"message-board/router"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	r := router.NewRouter(e)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	r.Init(httpPort)
}
