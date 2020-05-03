package database

import (
	"github.com/iborg-ai/core/src/models"
	"log"
)

// AutoMigrate will start mysql db migration
func AutoMigrate() {
	log.Println("Starting MySQL migrations")
	DBConn.AutoMigrate(&models.User{})
	log.Println("MySQL migrations complete")
}
