package database

import (
	"log"

	"github.com/iborg-ai/core/src/models"
)

// AutoMigrate will start mysql db migration
func AutoMigrate() {
	log.Println("Starting MySQL migrations")
	DBConn.AutoMigrate(&models.User{})
	log.Println("MySQL migrations complete")
}
