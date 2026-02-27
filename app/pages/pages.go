// aqui vai ser so aparte que vai renderizar os htmls da views
package pages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/app/models"
	"webapp/app/request"
	"webapp/app/response"
	utils "webapp/app/utils"
)

func LoadPageCreateUser(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "cadastro.html", nil)
}

func LoadHomePage(w http.ResponseWriter, req *http.Request) {
	url := "https://devbook-zqaw.onrender.com/publicacoes"
	resp, err := request.RequestAuth(req, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Erro ao buscar publicacoes: %v", err)
		utils.ExecuterTemplate(w, "home.html", nil)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.StatusCodeErr(w, resp)
		return
	}

	var posts []models.Posts
	if err = json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Err{Erro: err.Error()})
		return
	}

	fmt.Println(resp.StatusCode, err)
	utils.ExecuterTemplate(w, "home.html", posts)
}
