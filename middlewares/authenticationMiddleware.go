package middlewares

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"session-auth/models"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sessionToken := c.Value

		// search for session in database
		var sessions []models.Session

		content, err := ioutil.ReadFile(filepath.Join("db", "sessions.json"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(content, &sessions); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// search for session in sessions
		var userSession models.Session
		sessionFound := false
		for _, item := range sessions {
			if item.Value == sessionToken {
				userSession = item
				sessionFound = true
				break
			}
		}

		if !sessionFound {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("You are not logged in!"))
			return
		}

		if userSession.IsExpired() {
			// remove session from database
			var newSessions []models.Session

			for _, item := range sessions {
				if item.Value != sessionToken {
					newSessions = append(newSessions, item)
				}
			}

			jsonData, err := json.Marshal(newSessions)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := os.WriteFile(filepath.Join("db", "sessions.json"), jsonData, 0666); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// send response
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("You are not logged in!"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
