package main

import (
	"github.com/pravinba9495/iborg-core/src/database"
	"github.com/pravinba9495/iborg-core/src/router"
	"github.com/pravinba9495/iborg-core/src/server"
)

func main() {
	router.RegisterRoutes()
	database.ConnectDB()
	server.CreateServer()
}
