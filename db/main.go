package db

import (
	"context"
	"log"

	"github.com/rama-kairi/blog-api-golang-gin/ent"
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

// Ent Orm -
// https://entgo.io/docs/getting-started/
func InitEntDb() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
