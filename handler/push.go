package handler

import (
	"encoding/base64"
	"time"

	"github.com/gofiber/fiber/v2"
)

// DumpRequest represents the body of a new dump request
type DumpRequest struct {
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	Data      string    `json:"data"`
}

// Push handles /push requests
func Push(c *fiber.Ctx) error {
	dr := new(DumpRequest)
	if err := c.BodyParser(dr); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	d, err := base64.StdEncoding.DecodeString(dr.Data)
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	return c.Status(200).SendString(string(d))
}
