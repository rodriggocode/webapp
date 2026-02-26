// aqui vai ser so aparte que vai renderizar os htmls da views
package pages

import (
	"fmt"
	"log"
	"net/http"
	"webapp/app/request"
	utils "webapp/app/utils"
)

func LoadPageCreateUser(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "cadastro.html", nil)
}

func LoadHomePage(w http.ResponseWriter, req *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", "https://devbook-zqaw.onrender.com/publicacoes")
	resp, err := request.RequestAuth(req, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Erro ao buscar publicacoes: %v", err)
		utils.ExecuterTemplate(w, "home.hml", nil)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, err)
	utils.ExecuterTemplate(w, "home.html", nil)
}
