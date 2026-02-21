package users

import (
	"net/http"
	handlers "webapp/app/handlers/users"
	pages "webapp/app/pages"
)

func RouterUsers(router *http.ServeMux) {
	router.HandleFunc("/create/user", handlers.CreateUser)            // api
	router.HandleFunc("/cadastrar/usuario", pages.LoadPageCreateUser) // mostra o html
}
