package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL Dialect for GORM
	"github.com/qor/validations"
)

// ConnectDB instantiates a mongoDB connection
func ConnectDB() {
	var err error
	dbHost := os.Getenv("dbHost")
	if dbHost == "" {
		dbHost = "localhost"
	}
	DBConn, err = gorm.Open("mysql", "iborg:iborg@("+dbHost+")/iborg?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err.Error())
	}
	validations.RegisterCallbacks(DBConn)
	log.Println("Connected to database")
}
