package main

import (
	"fmt"

	"github.com/SoulKyu/go-website/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Println("Hello World !")
	app.Get("/", routes.Index)
	app.Listen(":8888")
}
