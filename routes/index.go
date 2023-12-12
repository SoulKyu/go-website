package routes

import (
	"fmt"

	"github.com/SoulKyu/go-website/dashboard"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	tmpl, err := dashboard.GetTemplate("index")
	if err != nil {
		fmt.Println("Error getting template:", err)
		return err
	}
	return dashboard.WriteTemplate(tmpl, nil, c)
}

func Souvenirs(c *fiber.Ctx) error {
	tmpl, err := dashboard.GetTemplate("souvenirs")
	if err != nil {
		fmt.Println("Error getting template", err)
		return err
	}
	souvenirs := generateSouvenirs()
	return dashboard.WriteTemplate(tmpl, fiber.Map{
		"Souvenirs": souvenirs,
	}, c)
}

func ServerAssets(c *fiber.Ctx) error {
	name := c.Params("name")
	return serveAssets(c, name)
}

func Submit(c *fiber.Ctx) error {
	userInput := c.FormValue("userInput")
	response := ""

	if userInput == "18/12/2020" {
		html := `<script>window.location.href = '/souvenirs';</script>`
		return c.SendString(html)
	}
	if userInput == "Taiga" || userInput == "taiga" {
		response = "Oh mon gros bb Taigounette d'amour... Mais non, c'est pas ca!"
	} else {
		response = "Ce n'est pas la bonne réponse, essaie encore <3"
	}

	return c.SendString(response)
}

func serveAssets(c *fiber.Ctx, name string) error {
	file, err := dashboard.AssetsFS.ReadFile(fmt.Sprintf("assets/%s.jpg", name))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Type("jpg").Send(file)
}
