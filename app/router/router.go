package router

import (
	"net/http"
	routes "webapp/app/router/routes/login"
)

func Gerar() *http.ServeMux {
	router := http.NewServeMux()
	routes.LoginRouter(router)
	return router
}
