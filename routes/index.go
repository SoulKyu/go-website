package routes

import (
	"embed"
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

func SouvenirHome(c *fiber.Ctx) error {
	name := c.Params("name")
	tmpl, err := dashboard.GetTemplate(name, "souvenir")
	if err != nil {
		fmt.Println("Error getting template", err)
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

func CorrezeSubmit(c *fiber.Ctx) error {
	userInput := c.FormValue("userInput")
	response := ""

	if userInput == "Puy Mary" {
		html := `
		<body class="flex items-center justify-center h-screen bg-gray-200">
		<p class="text-xl text-gray-800 font-bold mb-4">Bravo tu as trouvée !</p>
		<div class="flex">
			<div class="panel w-24 h-96" style="background-image: url('http://127.0.0.1:8888/IMG_5845.jpg')"></div>
			<div class="panel w-24 h-96" style="background-image: url('http://127.0.0.1:8888/IMG_5847.jpg')"></div>
			<div class="panel w-24 h-96" style="background-image: url('http://127.0.0.1:8888/IMG_5857.jpg')"></div>
			<div class="panel w-24 h-96" style="background-image: url('http://127.0.0.1:8888/IMG_5871.jpg')"></div>
			<div class="panel w-24 h-96" style="background-image: url('http://127.0.0.1:8888/IMG_5899.jpg')"></div>
		</div>
	
		<script>
			document.addEventListener('DOMContentLoaded', () => {
				const panels = document.querySelectorAll('.panel');
				panels.forEach(panel => {
					panel.addEventListener('click', () => {
						removeActiveClasses();
						panel.classList.add('active');
					});
				});
	
				function removeActiveClasses() {
					panels.forEach(panel => {
						panel.classList.remove('active');
					});
				}
			});
		</script>
	</body>
		`
		return c.SendString(html)
	}
	if userInput == "Taiga" || userInput == "taiga" {
		response = "Oh mon gros bb Taigounette d'amour... Mais non, c'est pas ca!"
		return c.SendString(response)
	} else {
		response = "Ce n'est pas la bonne réponse, essaie encore <3"
		return c.SendString(response)
	}
}

func CiskLemon(c *fiber.Ctx) error {
	userInput := c.FormValue("userInput")
	response := ""

	if userInput == "Lemon Cisk" || userInput == "Cisk Lemon" {
		html := `
		<body class="flex items-center justify-center h-screen bg-gray-200">
			<h1>Vive Malte</h1>
			<img src='http://127.0.0.1:8888/IMG_5018.jpg' alt='Funny Image' />
			<p>Surprise <3</p>
		</body>
		`
		return c.SendString(html)
	}
	if userInput == "Taiga" || userInput == "taiga" {
		response = "Oh mon gros bb Taigounette d'amour... Mais non, c'est pas ca!"
		return c.SendString(response)
	} else {
		response = "Petit indice ?, CITRON !"
		return c.SendString(response)
	}
}

func SetupStatic(app *fiber.App, root string, fs embed.FS) {
	app.Use(root, func(c *fiber.Ctx) error {
		file, err := fs.ReadFile("assets" + c.OriginalURL())
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Send(file)
	})
}
