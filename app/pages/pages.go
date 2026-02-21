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
	utils.ExecuterTemplate(w, "home.html", nil)
}
