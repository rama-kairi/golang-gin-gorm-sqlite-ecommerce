package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitGormDb() {
	var err error
	dbDsn := viper.GetString("DB_DSN")
	Db, err = gorm.Open(sqlite.Open(dbDsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
