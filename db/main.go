package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitGormDb() {
	var err error
	dbDsn := os.Getenv("DB_DSN")
	Db, err = gorm.Open(sqlite.Open(dbDsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
