package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL Dialect for GORM
)

// ConnectDB instantiates a mongoDB connection
func ConnectDB() {
	var err error
	DBConn, err = gorm.Open("mysql", "iborg:iborg@(192.168.1.116)/iborg?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Connected to database")
}
