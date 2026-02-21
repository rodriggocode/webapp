// aqui vai ser so aparte que vai renderizar os htmls da views
package pages

import (
	// Keep this import

	"net/http"
	utils "webapp/app/utils"
)

func LoadPageCreateUser(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "cadastro.html", nil)
}

func LoadHomePage(w http.ResponseWriter, req *http.Request) {
	/*url := fmt.Sprintf("%s/posts", config.APIURL)
	res, err := request.RequestAuth(req, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Err{Erro: err.Error()})
		return
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		response.StatusCodeErr(w, res) // Keep local's correct var name
		return
	}

	var posts []models.Posts
	if err = json.NewDecoder(res.Body).Decode(&posts); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Err{Erro: err.Error()})
		return
		}*/

	utils.ExecuterTemplate(w, "home.html", nil)
}
