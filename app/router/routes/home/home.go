package home

import (
	"net/http"
	handlers "webapp/app/handlers/home"
	"webapp/app/middlewares"
	page "webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
	router.HandleFunc("/home-handler", handlers.HomePage)                               // carrega minha logica com o back
	router.HandleFunc("/home", middlewares.Logger(middlewares.Auth(page.LoadHomePage))) // carrega minha pagina de home.html
}
