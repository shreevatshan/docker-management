package main

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func loadingGIF(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join(currentdir, "loading.gif")))
	tmpl.Execute(w, nil)
}
