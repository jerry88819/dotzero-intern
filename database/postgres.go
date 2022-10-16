package postgres

import (
	"fmt"
	"log"
	"time"

	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "log"
	"os"
	// "time"
)

func Init() (db *gorm.DB, err error) {

	// db, err = gorm.Open("postgres", "postgresql://dz_postgres@pgdb-dz:wkMYwZQs$211!@pgdb-dz.postgres.database.azure.com/dev_tenant?sslmode=require")
	// db, err = gorm.Open("postgres", "postgresql://jerry:EpE39I54a58qS@35.221.198.8/dev_tenant?sslmode=require")

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	pgUrl := os.Getenv("PG_URL")
	fmt.Println(pgUrl)
	// pgUrl := "postgresql://jerry:EpE39I54a58qS@34.80.53.176/dev_tenant?sslmode=require"
	db, err = gorm.Open(postgres.Open(pgUrl), &gorm.Config{})
	if err != nil {
		log.Println("Open pg db error:", err)
		return
	}

	_db, err := db.DB()
	if err != nil {
		log.Println("Set db connection pool failed:", err)
		return
	}

	_db.SetMaxIdleConns(2)
	_db.SetMaxOpenConns(5)
	_db.SetConnMaxLifetime(time.Hour)

	return
} // Init()
