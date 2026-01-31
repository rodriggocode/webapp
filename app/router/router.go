package router

import (
	"net/http"
	routes "webapp/app/router/routes/login"
)

func Gerar() *http.ServeMux {
	router := http.NewServeMux()
	routes.LoginRouter(router)

	/*  aqui serve para os aquivos staticos css e js.
	ainda nao sei se ele fica na rota aqui a principal ou em cada uma separada e depois eu vou ver
	mas tmb nao sei se preciso ter isso pra funcionar
	agora eu vou tirar e depois eu coloco, caso veja que ele e preciso
	*/
	fileServer := http.FileServer(http.Dir("./assets"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	return router
}
