package main

import (
	"log"
	"message-board/infrastructure/postgres"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	db, dbClose, err := postgres.Connect()
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer dbClose()

	r := InitRouter(e, db)
	r.Init("8080")
}
