package login

import (
	"net/http"
	handlers "webapp/app/handlers/login"
)

func LoginRouter(router *http.ServeMux) {
	router.HandleFunc("/", handlers.Login)
	router.HandleFunc("/login", handlers.Login)
}
