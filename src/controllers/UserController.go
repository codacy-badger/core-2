package controllers

import (
	"net/http"

	"github.com/gofiber/fiber"

	"github.com/iborg-ai/core/src/middleware"
	"github.com/iborg-ai/core/src/models"
	"github.com/iborg-ai/core/src/services"
)

// UserController represents the entry point for the User API
func UserController(c *fiber.Ctx) {
	middleware.AddHeaders(c)
	if c.Method() != "OPTIONS" {
		switch c.Method() {
		case "POST":
			saveUser(c)
			break
		default:
			c.Status(http.StatusMethodNotAllowed).JSON(&models.HTTPErrorStatus{Status: http.StatusMethodNotAllowed, Message: http.StatusText(http.StatusMethodNotAllowed)})
			break
		}
	} else {
		c.SendStatus(http.StatusOK)
	}

}

func saveUser(c *fiber.Ctx) {

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(http.StatusBadRequest).JSON(&models.HTTPErrorStatus{Status: http.StatusBadRequest, Message: http.StatusText(http.StatusBadRequest)})
		return
	}

	result := services.SaveUser(user)
	c.Status(result.Status).JSON(result)
}
