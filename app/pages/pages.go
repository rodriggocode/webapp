// aqui vai ser so aparte que vai renderizar os htmls da views
package pages

import (
	// Keep this import

	"fmt"
	"net/http"
	"webapp/app/config"
	"webapp/app/request"
	utils "webapp/app/utils"
)

func LoadPageCreateUser(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "cadastro.html", nil)
}

func LoadHomePage(w http.ResponseWriter, req *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	resp, err := request.RequestAuth(req, http.MethodGet, url, nil)
	fmt.Println(resp.StatusCode, err)
	utils.ExecuterTemplate(w, "home.html", nil)
}
