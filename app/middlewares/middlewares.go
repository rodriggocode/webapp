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
/*func Auth(nextFunction http.HandlerFunc) http.HandlerFunc {
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
}*/

func Auth(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		token, err := cookies.ReadToken(req)
		if err != nil {
			log.Printf("Erro ao ler cookie: %v", err)
			http.Redirect(w, req, "/logar", 302)
			return
		}
		log.Printf("Token encontrado: %s", token[:20]) // só os primeiros 20 caracteres

		apiReq, _ := http.NewRequest("GET", "https://devbook-zqaw.onrender.com/publicacoes", nil)
		apiReq.Header.Set("Authorization", "Bearer "+token)
		apiReq.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(apiReq)
		if err != nil {
			log.Printf("Erro ao chamar API: %v", err)
			http.Redirect(w, req, "/logar", 303)
			return
		}
		log.Printf("Status da API: %d", resp.StatusCode)
		if resp.StatusCode != http.StatusOK {
			http.Redirect(w, req, "/logar", 303)
			return
		}

		nextFunction(w, req)
	}
}
