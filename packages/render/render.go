package render

import (
	"fmt"
	"net/http"
	helpers "nu3w4v3/packages/helpers"
	"text/template"
)

var tmplCache = make(map[string]*template.Template)

func RenderTemplate(writer http.ResponseWriter, request *http.Request, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tmplCache[t]
	if !inMap {
		err = makeTemplateCache(t)
		helpers.ErrorCheck(err)
	} else {
		fmt.Println("Template found in cache")
	}

	tmpl = tmplCache[t]
	err = tmpl.Execute(writer, request)
	helpers.ErrorCheck(err)
}

func makeTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	helpers.ErrorCheck(err)

	tmplCache[t] = tmpl

	return nil
}
