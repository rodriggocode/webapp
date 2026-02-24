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
		token, err := cookies.ReadToken(req)
		if err != nil {
			http.Redirect(w, req, "/logar", 302)
			return
		}
		apiReq, _ := http.NewRequest("GET", "https://devbook-zqaw.onrender.com/publicacoes", nil)
		apiReq.Header.Set("Authorization", "Bearer "+token)
		apiReq.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(apiReq)
		if err != nil || resp.StatusCode != http.StatusOK {
			http.Redirect(w, req, "/logar", 303)
			return
		}
		nextFunction(w, req)
	}
}
