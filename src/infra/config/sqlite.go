package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(path string) *gorm.DB{
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Panicf("error connect sqlite - %+v", err)
	}

	return db
}