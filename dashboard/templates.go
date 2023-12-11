package dashboard

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

const (
	HeadTemplateName   = "head.gohtml"
	FooterTemplateName = "footer.gohtml"
	extension          = "/*.html"
)

var (
	//go:embed templates/*.html
	TemplateFS               embed.FS
	templates                map[string]*template.Template
	defaultIncludedTemplates = []string{
		"head",
		"footer",
	}
)

type baseTemplateData struct {
	BasePath string
	Data     interface{}
	JSON     template.JS
}

func GetTemplate(name string) (*template.Template, error) {
	tmpl := template.New(name)

	// Parse the specific template file along with header and footer
	_, err := tmpl.ParseFS(TemplateFS, "templates/"+name+".html", "templates/header.html", "templates/footer.html")
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func WriteTemplate(tmpl *template.Template, data interface{}, c *fiber.Ctx) error {
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.SendString(err.Error())
		return err
	}

	c.Status(fiber.StatusOK)
	return c.SendString(buf.String())
}

func DebugTemplate(tmpl *template.Template, data interface{}) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		fmt.Println("Error executing template:", err)
	} else {
		fmt.Println("Template output:", buf.String())
	}
}
