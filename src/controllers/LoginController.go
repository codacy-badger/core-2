package controllers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/iborg-ai/core/src/middleware"
	"github.com/iborg-ai/core/src/models"
	"github.com/iborg-ai/core/src/services"
)

// LoginController represents the entry point for the Login API
func LoginController(c *fiber.Ctx) {
	log.Println("Login API")
	middleware.AddHeaders(c)

	switch c.Method() {
	case "POST":
		doLogin(c)
		break
	default:
		c.SendStatus(http.StatusMethodNotAllowed)
		break
	}
}

func doLogin(c *fiber.Ctx) {
	log.Println("Do Login")

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return
	}

	result := services.DoLogin(user.Email, user.Password)
	c.Status(result.Status).JSON(result)
}
