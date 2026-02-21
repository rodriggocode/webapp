package home

import (
	"net/http"
	handlers "webapp/app/handlers/home"
	"webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
	router.HandleFunc("/homepage", handlers.Home) // vem da api
	router.HandleFunc("home", pages.LoadHomePage) // html front
}
