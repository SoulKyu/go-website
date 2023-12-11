package routes

import (
	"fmt"

	"github.com/SoulKyu/go-website/dashboard"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	tmpl, err := dashboard.GetTemplate("test")
	if err != nil {
		fmt.Println("Error getting template:", err)
	} else {
		dashboard.DebugTemplate(tmpl, nil)
	}
	if err != nil {
		return err
	}
	return dashboard.WriteTemplate(tmpl, nil, c)
}
