package database

import (
	"golang-mock/models"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	user := ""
	password := ""
	server := ""
	database := ""

	connectionString := "sqlserver://" + user + ":" + password + "@" + server + "?database=" + database
	DB, err = gorm.Open(sqlserver.Open(connectionString))

	if err != nil {
		log.Panic("Error connecting to database")
	}

	DB.AutoMigrate(&models.Student{})
}
