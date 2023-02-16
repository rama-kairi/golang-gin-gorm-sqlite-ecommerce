package db

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
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
	dsn := "host=localhost user=postgres password=postgres dbname=ecomm port=5432 sslmode=disable"
	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
