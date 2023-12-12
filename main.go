package main

import (
	"fmt"

	"github.com/SoulKyu/go-website/dashboard"
	"github.com/SoulKyu/go-website/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Println("Hello World !")
	app.Get("/", routes.Index)
	app.Post("/submit", routes.Submit)
	app.Get("/souvenirs", routes.Souvenirs)
	app.Get("/souvenir/:name", routes.SouvenirHome)
	routes.SetupStatic(app, "/", dashboard.AssetsFS)
	app.Listen(":8888")
}
