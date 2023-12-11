package routes

import (
	"fmt"

	"github.com/SoulKyu/go-website/dashboard"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	tmpl, err := dashboard.GetTemplate("index")
	fmt.Println(tmpl.Templates())
	if err != nil {
		return err
	}
	return dashboard.WriteTemplate(tmpl, nil, c)
}
