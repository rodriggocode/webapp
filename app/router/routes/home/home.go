package home

import (
	"net/http"
	handlers "webapp/app/handlers/home"
	"webapp/app/middlewares"
	"webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
	router.HandleFunc("/logar", handlers.Home)                                            // vem da api
	router.HandleFunc("/login", middlewares.Logger(middlewares.Auth(pages.LoadHomePage))) // html front
}
