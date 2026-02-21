package home

import (
	"net/http"
	handlers "webapp/app/handlers/home"
	"webapp/app/middlewares"
	page "webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
	router.HandleFunc("/home", handlers.HomePage)                                            // carrega minha logica com o back
	router.HandleFunc("/home-page", middlewares.Logger(middlewares.Auth(page.LoadHomePage))) // carrega minha pagina de home.html
}
