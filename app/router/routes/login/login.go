package login

import (
	"net/http"
	handlers "webapp/app/handlers/login"
)

func LoginRouter(router *http.ServeMux) {
	router.HandleFunc("/", handlers.Login)            // aqui mostra a tela de login
	router.HandleFunc("/publicacoes", handlers.Login) // aqui mostra a tela de login tmb
	router.HandleFunc("/home", handlers.PostLogin)    // aqui tem que ficar /login da url da api
}
