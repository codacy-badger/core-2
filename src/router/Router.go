package router

import (
	"github.com/gofiber/fiber"
	"github.com/iborg-ai/core/src/controllers"
)

// RegisterRoutes registers all API routes for the application
func RegisterRoutes(app *fiber.App) {
	app.Group("/api/users", controllers.UserController)
	app.Group("/api/login", controllers.LoginController)
}
