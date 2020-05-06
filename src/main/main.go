package main

import (
	"github.com/iborg-ai/core/src/database"
	"github.com/iborg-ai/core/src/server"
)

func main() {
	database.ConnectDB()
	database.AutoMigrate()
	server.CreateServer()
}
