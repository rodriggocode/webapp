package home

import (
	"net/http"
	"webapp/app/middlewares"
	"webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
	router.HandleFunc("/home", middlewares.Logger(middlewares.Auth(pages.LoadHomePage))) // carrega minha pagina de home.html
	// aqui vai ficar sem o handler so para eu fazer um teste
}
