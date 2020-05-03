package database

import (
	"github.com/pravinba9495/iborg-core/src/models"
	"log"
)

// AutoMigrate will start mysql db migration
func AutoMigrate() {
	log.Println("Starting MySQL migrations")
	DBConn.AutoMigrate(&models.User{})
	log.Println("MySQL migrations complete")
}
