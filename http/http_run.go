package http

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter() {
	app := fiber.New()

	app.Get("/")
}
