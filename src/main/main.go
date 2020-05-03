package main

import (
	"github.com/iborg-ai/core/src/database"
	"github.com/iborg-ai/core/src/router"
	"github.com/iborg-ai/core/src/server"
)

func main() {
	router.RegisterRoutes()
	database.ConnectDB()
	database.AutoMigrate()
	server.CreateServer()
}