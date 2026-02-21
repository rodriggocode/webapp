package home

import (
	"net/http"
	handlers "webapp/app/handlers/home"
	"webapp/app/middlewares"
	page "webapp/app/pages"
)

func HomeRouter(router *http.ServeMux) {
<<<<<<< HEAD
	router.HandleFunc("/home", handlers.HomePage)                                            // carrega minha logica com o back
	router.HandleFunc("/home-page", middlewares.Logger(middlewares.Auth(page.LoadHomePage))) // carrega minha pagina de home.html
=======
	//router.HandleFunc("/home", handlers.HomePage)                                            // carrega minha logica com o back
	router.HandleFunc("/home", handlers.HomePage) // carrega minha pagina de home.html

	router.HandleFunc("/home-page", middlewares.Logger(middlewares.Auth(page.LoadHomePage)))
>>>>>>> 012b24e24580126d4d8b497759896eead02de80f
}
