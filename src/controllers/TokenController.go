package controllers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/iborg-ai/core/src/middleware"
	"github.com/iborg-ai/core/src/models"
)

// TokenController represents the entry point for the Token API
func TokenController(c *fiber.Ctx) {
	log.Println("Token API")
	middleware.AddHeaders(c)

	switch c.Method() {
	case "POST":
		validateToken(c)
		break
	case "OPTIONS":
		c.SendStatus(200)
		break
	default:
		c.Status(http.StatusMethodNotAllowed).JSON(&models.HTTPErrorStatus{Status: http.StatusMethodNotAllowed, Message: http.StatusText(http.StatusMethodNotAllowed)})
		break
	}
}

func validateToken(c *fiber.Ctx) {
	log.Println("Do Validate Token")

	token := new(models.TokenPayload)
	if err := c.BodyParser(token); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return
	}
	status := middleware.ValidateToken(token.Token)
	c.Status(status).JSON(&models.HTTPErrorStatus{Status: status, Message: http.StatusText(status)})
}
