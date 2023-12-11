package dashboard

import (
	"embed"
	"html/template"
)

const (
	HeadTemplateName   = "head.gohtml"
	FooterTemplateName = "footer.gohtml"
	extension          = "/*.html"
)

var (
	files                    embed.FS
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
