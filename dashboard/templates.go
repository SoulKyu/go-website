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

	_, err := tmpl.ParseFS(TemplateFS, fmt.Sprintf("templates/%s.html", name), "templates/header.html", "templates/footer.html")
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

	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

	c.Status(fiber.StatusOK)
	return c.SendString(buf.String())
}
