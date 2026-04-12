package post

import (
	"net/http"
	handler "webapp/app/handlers/post"
	"webapp/app/middlewares"
)

func RouterPost(router *http.ServeMux) {
	router.HandleFunc("/create/post", middlewares.Auth(handler.CreatePost))
}
