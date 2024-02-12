package main

import (
	"net/http"
	"text/template"
)

func loadingGIF(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("loading.gif"))
	tmpl.Execute(w, nil)
}
