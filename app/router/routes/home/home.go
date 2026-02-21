package home

import (
	"net/http"
	"webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
	router.HandleFunc("/home", pages.LoadHomePage) // carrega minha pagina de home.html
}
