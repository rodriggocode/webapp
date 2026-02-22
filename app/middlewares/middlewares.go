package middlewares

import (
	"log"
	"net/http"
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

		// faz a requisição ao backend passando o cookie
		apiReq, _ := http.NewRequest("GET", "https://devbook-zqaw.onrender.com/publicacoes", nil)

		// repassa o cookie que o navegador enviou
		apiReq.Header.Set("Cookie", req.Header.Get("Cookie"))
		apiReq.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(apiReq)

		if err != nil || resp.StatusCode != http.StatusOK {
			http.Redirect(w, req, "/logar", 302)
			return
		}

		nextFunction(w, req)
	}
}
