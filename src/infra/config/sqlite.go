package config

import (
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB{

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Silent, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  Colorful:                  false,          // Disable color
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