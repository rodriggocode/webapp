package middlewares

import (
	"log"
	"net/http"
	"webapp/app/cookies"
)

// escreve infomacoes da requisicao no termianl
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("\n %s %s %s", req.Method, req.RequestURI, req.Host)
		nextFunction(w, req)
	}
}

// verifica a existencia de cookies
func Auth(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		_, err := cookies.ReadToken(req)
		if err != nil {
			http.Error(w, "Nao autorizado", http.StatusUnauthorized)
			return
		}
		nextFunction(w, req)
	}
}
