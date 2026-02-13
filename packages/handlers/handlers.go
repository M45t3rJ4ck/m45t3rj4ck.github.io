package handlers

import (
	"net/http"
	render "nu3w4v3/packages/render"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "home.page.tmpl")
}

func AboutHandler(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "about.page.tmpl")
}

func ContactMeHandler(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "contactMe.page.tmpl")
}

func CssHandler(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "css.page.tmpl")
}

func HtmlHandler(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "html.page.tmpl")
}

func MoreHelpHandler(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "moreHelp.page.tmpl")
}
