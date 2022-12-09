package controllers

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func ShowPage(w http.ResponseWriter, r *http.Request) {
	var filename string
	path := r.URL.Path
	switch path {
	case "/":
		filename = "index.html"
	case "/admin":
		filename = "admin.html"
	case "/register":
		filename = "register.html"
	case "/login":
		filename = "login.html"
	}

	tmpl, err := template.ParseFiles(filepath.Join("views", filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
