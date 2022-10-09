package main

import (
	"context"
	"message-board/infrastructure/postgres"
)

func main() {
	db, dbClose, err := postgres.Connect()
	if err != nil {
		panic(err)
	}
	defer dbClose()

	ctx := context.Background()
	conn := db.Conn(ctx)
	conn.Table("users").AutoMigrate(&postgres.UserDTO{})
}
