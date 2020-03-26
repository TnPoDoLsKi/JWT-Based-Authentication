package server

import (
	"JWT-Based-Authentication/controllers"
	"net/http"
)

func Run(port string) {

	userRoutes()
	http.ListenAndServe(":8000", nil)
}

func userRoutes() {
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/resgister", controllers.Register)
}
