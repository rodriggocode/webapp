package users

import (
	"net/http"
	handlers "webapp/app/handlers/users"
)

func RouterUsers(router *http.ServeMux) {
	router.HandleFunc("/cadastrar/usuario", handlers.CreateUsers)
}
