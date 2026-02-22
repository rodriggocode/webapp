package home

import (
	"net/http"
	handlers "webapp/app/handlers/home"
	"webapp/app/middlewares"
	"webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
	router.HandleFunc("/publicacoes", handlers.Home)     // vem da api
	router.HandleFunc("/home", middlewares.Logger(middlewares.Auth(pages.LoadHomePage)) // html front
}
