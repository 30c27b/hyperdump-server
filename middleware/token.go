package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

const token = "example.token"

// THIS IS FOR TESTING ONLY

// Auth is a middleware to check for a specific token
func Auth(c *fiber.Ctx) error {
	h := c.Get("Authorisation")
	p := strings.TrimPrefix(h, "Bearer ")
	if h == p {
		return c.Status(401).SendString("Malformed Authorisation header.")
	}
	if p != token {
		return c.Status(401).SendString("Invalid token.")
	}
	return c.Next()
}
