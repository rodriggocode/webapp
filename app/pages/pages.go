// aqui vai ser so aparte que vai renderizar os htmls da views
package pages

import (
	"net/http"
	utils "webapp/app/utils"
)

func LoadPageCreateUser(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "cadastro.html", nil)
}
