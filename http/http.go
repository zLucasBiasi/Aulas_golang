package main

import (
	"net/http"
	"text/template"
)

var templates *template.Template

type user struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := user{"Biasi", "emaildobiasi@gmail.com"}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	http.ListenAndServe(":5000", nil)

}
