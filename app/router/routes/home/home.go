package home

import (
	"net/http"
	handlers "webapp/app/handlers/home"
	"webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
	router.HandleFunc("/home", handlers.Home)          // vem da api
	router.HandleFunc("home-page", pages.LoadHomePage) // html front
}
