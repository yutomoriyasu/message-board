package main

import (
	"message-board/router"

	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	e := echo.New()

	r := router.NewRouter(e)

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		"root",
		"password",
		"test",
		"db",
		"5432",
	)
	db, _ := gorm.Open(postgres.Open(dsn))

	fmt.Println(db.Statement.Vars...)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	r.Init(httpPort)
}
