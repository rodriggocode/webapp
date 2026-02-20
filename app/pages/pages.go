// aqui vai ser so aparte que vai renderizar os htmls da views
package pages

import (
	"encoding/json" // Keep this import
	"fmt"
	"net/http"
	"webapp/app/config"
	"webapp/app/models"
	"webapp/app/request"
	"webapp/app/response"
	utils "webapp/app/utils"
)

func LoadPageCreateUser(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "cadastro.html", nil)
}

func LoadHomePage(w http.ResponseWriter) {
	url := fmt.Sprintf("%s/posts", config.APIURL) // Use remote's URL path, keep 'res' var name
	res, err := request.RequestAuth(w, http.MethodGet, url, nil) // Adjusted for RequestAuth signature
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Err{Erro: err.Error()})
		return
	}
	defer res.Body.Close() // Keep local's correct defer
	if res.StatusCode >= 400 {
		response.StatusCodeErr(w, res) // Keep local's correct var name
		return
	}

	var posts []models.Posts
	if err = json.NewDecoder(res.Body).Decode(&posts); err != nil { // Keep local's logic for decoding posts
		response.JSON(w, http.StatusUnprocessableEntity, response.Err{Erro: err.Error()})
		return
	}

	utils.ExecuterTemplate(w, "home.html", nil)
}