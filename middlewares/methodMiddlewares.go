package middlewares

import "net/http"

func RequireGET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequirePOST(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
