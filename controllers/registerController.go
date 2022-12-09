package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"session-auth/models"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	// reading the request
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	creds := models.Credentials{
		Username: username,
		Password: password,
	}

	// read the database
	var users []models.Credentials

	content, err := ioutil.ReadFile(filepath.Join("db", "users.json"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(content, &users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// append the new value
	users = append(users, creds)

	// write the database
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile(filepath.Join("db", "users.json"), jsonData, 0666); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Register success!"))
}
