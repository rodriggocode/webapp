package home

import (
	"net/http"
	"webapp/app/middlewares"
	page "webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
	router.HandleFunc("/home-page", middlewares.Logger(middlewares.Auth(page.LoadHomePage))) // carrega minha logica com o back
	//router.HandleFunc("/home", middlewares.Logger(middlewares.Auth(page.LoadHomePage))) // carrega minha pagina de home.html
}
