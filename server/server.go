package server

import (
	"JWT-Based-Authentication/controllers"
	"log"
	"net/http"
)

func Run(port string) {

	userRoutes()
	log.Print(http.ListenAndServe(":8000", nil))
}

func userRoutes() {
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/resgister", controllers.Register)
}
