package database

import (
	"fmt"
	"log"
	"w8s/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB(configs *models.Config) {

	configDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", configs.Database.Host, configs.Database.Port, configs.Database.User, configs.Database.Pass, configs.Database.Name)

	db, err = gorm.Open(postgres.Open(configDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	db.Debug().AutoMigrate(&models.Person{}, &models.Order{}, &models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
