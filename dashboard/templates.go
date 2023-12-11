package dashboard

import (
	"bytes"
	"embed"
	"html/template"

	"github.com/gofiber/fiber"
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

	// Parse the specific template file
	_, err := tmpl.ParseFS(TemplateFS, "templates/"+name+".html")
	if err != nil {
		return nil, err
	}

	// Parse other templates that should be included by default
	_, err = tmpl.ParseFS(TemplateFS, "templates/header.html", "templates/footer.html")
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func WriteTemplate(tmpl *template.Template, data interface{}, c *fiber.Ctx) {
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		c.SendString(err.Error())
	}

	c.Status(fiber.StatusOK)
	c.SendString(buf.String())
}
