package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	
	if err := godotenv.Load("../.env"); err != nil {
		panic("CONFIGURATION ERROR: .ENV")
	}

	path := os.Getenv("PATH_SQLITE")

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicf("error connect sqlite - %+v", err)
	}

	return db
}
