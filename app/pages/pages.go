// aqui vai ser so aparte que vai renderizar os htmls da views
package pages

import (
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

func LoadHomePage(w http.ResponseWriter, req *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	respo, err := request.RequestAuth(req, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Err{Erro: err.Error()})
		return
	}
	defer req.Body.Close()
	if respo.StatusCode >= 400 {
		response.StatusCodeErr(w, respo)
		return
	}

	var posts []models.Posts
	
	utils.ExecuterTemplate(w, "home.html", nil)
}
