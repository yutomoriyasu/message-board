package main

import (
	"fmt"
	"log"
	"message-board/infrastructure/postgres"
	"message-board/router"

	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	r := router.NewRouter(e)

	db, dbClose, err := postgres.Connect()
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer dbClose()
	fmt.Println(db)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	r.Init(httpPort)
}
