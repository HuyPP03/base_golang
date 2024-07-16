package database

import (
	"log"

	"github.com/HuyPP03/learn/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	log.Println("Successfully connected to the database!")

	DB = database
}
