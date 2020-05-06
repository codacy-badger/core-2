package server

import (
	"os"

	"github.com/gofiber/fiber"
	"github.com/iborg-ai/core/src/router"
)

// CreateServer creates a HTTP server
func CreateServer() {
	port := os.Getenv("IBORG_PORT")
	if port == "" {
		port = "8080"
	}
	app := fiber.New()
	router.RegisterRoutes(app)
	app.Listen(port)
}
