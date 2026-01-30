package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadingTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecuterTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
