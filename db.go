package main

import (
	m "butaforia.io/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func MakeMigrations() {

	err := DB.AutoMigrate(
		&m.User{},
	)

	if err != nil {
		info := map[string]interface{}{
			"func":  "ApplyMigrations(models)",
			"error": err,
		}
		log.Fatal(info)
	}

}
