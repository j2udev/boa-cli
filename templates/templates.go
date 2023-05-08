package templates

import (
	"embed"
	"html/template"

	"github.com/charmbracelet/log"
)

//go:embed *.tmpl
var templates embed.FS

func NewCobraCmdTemplate() *template.Template {
	tmpl, err := template.New("cobraCmd").Funcs(template.FuncMap{
		"pascalCase": pascalCase,
		"lastIndex":  lastIndex,
		"moduleName": moduleName,
	}).ParseFS(templates, "command.tmpl")
	if err != nil {
		log.Error(err)
	}
	return tmpl
}

func NewPkgTemplate() *template.Template {
	tmpl, err := template.New("pkg").Funcs(template.FuncMap{
		"pascalCase": pascalCase,
	}).ParseFS(templates, "pkg.tmpl")
	if err != nil {
		log.Error(err)
	}
	return tmpl
}

func NewMainTemplate() *template.Template {
	tmpl, err := template.New("main").Funcs(template.FuncMap{
		"pascalCase": pascalCase,
	}).ParseFS(templates, "main.tmpl")
	if err != nil {
		log.Error(err)
	}
	return tmpl
}

func NewBoaConfigTemplate() *template.Template {
	tmpl, err := template.New("boa.yml").ParseFS(templates, "boa.tmpl")
	if err != nil {
		log.Error(err)
	}
	return tmpl
}
