package main

import (
	"net/http"
	handlers "nu3w4v3/packages/handlers"
	helpers "nu3w4v3/packages/helpers"
)

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	helpers.ErrorCheck(err)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/contactMe", handlers.ContactMeHandler)
	http.HandleFunc("/css", handlers.CssHandler)
	http.HandleFunc("/html", handlers.HtmlHandler)
	http.HandleFunc("/moreHelp", handlers.MoreHelpHandler)

	err := http.ListenAndServe(":5000", nil)
	helpers.ErrorCheck(err)
}
