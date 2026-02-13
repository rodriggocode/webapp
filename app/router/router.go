package router

import (
	"net/http"
	router_home "webapp/app/router/routes/home"
	routes "webapp/app/router/routes/login"
	router_users "webapp/app/router/routes/users"
)

func Gerar() *http.ServeMux {
	router := http.NewServeMux()
	routes.LoginRouter(router)
	router_users.RouterUsers(router)
	router_home.HomeRouter(router)
	/*  aqui serve para os aquivos staticos css e js.
	ainda nao sei se ele fica na rota aqui a principal ou em cada uma separada e depois eu vou ver
	mas tmb nao sei se preciso ter isso pra funcionar
	agora eu vou tirar e depois eu coloco, caso veja que ele e preciso
	*/
	fileServer := http.FileServer(http.Dir("./assets"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	return router
}
