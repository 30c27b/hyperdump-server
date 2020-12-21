package main

import (
	"flag"

	"github.com/30c27b/dump-server/handler"
	"github.com/30c27b/dump-server/middleware"
	"github.com/gofiber/fiber/v2"
)

var (
	port = flag.String("port", "3000", "Listening port")
)

func main() {
	app := fiber.New()

	app.Put("/push", middleware.Auth, handler.Push)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(404).SendString("not found")
	})

	app.Listen(":" + *port)
}
