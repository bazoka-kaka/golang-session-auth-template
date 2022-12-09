package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"session-auth/models"
	"time"

	"github.com/google/uuid"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// read from the database
	var users []models.Credentials

	content, err := os.ReadFile(filepath.Join("db", "users.json"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(content, &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// search for username match
	found := false
	for _, user := range users {
		if user.Username == username && user.Password == password {
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Wrong password/username"))
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(time.Minute * 30)

	session := models.Session{
		Username: username,
		Expiry:   expiresAt,
		Value:    sessionToken,
	}

	// read sessions database
	var sessions []models.Session

	content, err = ioutil.ReadFile(filepath.Join("db", "sessions.json"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(content, &sessions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// append new session to database
	sessions = append(sessions, session)

	// write sessions to database
	jsonData, err := json.Marshal(sessions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(filepath.Join("db", "sessions.json"), jsonData, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})

	w.Write([]byte(fmt.Sprintf("Login Success!\n\nsession_token=%s\n\n", sessionToken)))
}
