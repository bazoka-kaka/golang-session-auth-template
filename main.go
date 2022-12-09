package main

import (
	"fmt"
	"net/http"
	"session-auth/controllers"
	"session-auth/middlewares"
)

func main() {
	// show pages
	http.Handle("/", middlewares.RequireGET(http.HandlerFunc(controllers.ShowPage)))
	http.Handle("/register", middlewares.RequireGET(http.HandlerFunc(controllers.ShowPage)))
	http.Handle("/login", middlewares.RequireGET(http.HandlerFunc(controllers.ShowPage)))
	http.Handle("/admin", middlewares.Authenticate(middlewares.RequireGET(http.HandlerFunc(controllers.ShowPage))))

	// handlers
	http.Handle("/user/register", middlewares.RequirePOST(http.HandlerFunc(controllers.HandleRegister)))
	http.Handle("/user/login", middlewares.RequirePOST(http.HandlerFunc(controllers.HandleLogin)))

	fmt.Println("server running on port 3000")
	http.ListenAndServe(":3000", nil)
}
