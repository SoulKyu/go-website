package routes

import (
	"github.com/SoulKyu/go-website/dashboard"
	"github.com/gofiber/fiber"
)

func Index(c *fiber.Ctx) {

	tmpl, _ := dashboard.GetTemplate("index")
	dashboard.WriteTemplate(tmpl, nil, c)
}
