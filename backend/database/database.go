package database

import (
	"eventhub/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "host=localhost port=5432 user=sakreda password=qweqweqwE5! dbname=eventhub sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}

	DB = db
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Event{})
}
