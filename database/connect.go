package database

import (
	"github/users/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
	database, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_CONNECTION_URL")), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}