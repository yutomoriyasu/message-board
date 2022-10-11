package postgres

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB interface {
	Conn(context.Context) *gorm.DB
}

type gormDB struct {
	db *gorm.DB
}

var currentConnection *gormDB = nil

func (db *gormDB) Conn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(&struct{}{}).(*gorm.DB)
	if !ok {
		return db.db.WithContext(ctx)
	}
	return tx.WithContext(ctx)
}

func Connect() (DB, func() error, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		"root",
		"password",
		"test",
		"db",
		"5432",
	)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, nil, err
	}

	if os.Getenv("APP_ENV") == "development" {
		db = db.Debug()
	}

	if currentConnection == nil {
		// when initial connect
		currentConnection = &gormDB{db}
	} else {
		// when reconnect
		sqlDB, err := currentConnection.db.DB()
		if err != nil {
			return nil, nil, err
		}
		sqlDB.Close()
		currentConnection.db = db
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	return currentConnection, sqlDB.Close, err
}
