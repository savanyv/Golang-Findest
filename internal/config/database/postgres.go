package database

import (
	"log"

	"github.com/savanyv/Golang-Findest/internal/config"
	"github.com/savanyv/Golang-Findest/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(c config.Config) (*gorm.DB, error) {
	dsn := "host=" + c.Database.PGHost + " user=" + c.Database.PGUser + " password=" + c.Database.PGPassword + " dbname=" + c.Database.PGDatabase + " port=" + c.Database.PGPort + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect database")
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Transaction{},
	); err != nil {
		log.Println("failed to migrate database")
		return nil, err
	}

	DB = db

	return db, nil
}