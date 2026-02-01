package users

import (
	"net/http"
	"webapp/app/utils"
)

func CreateUsers(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "cadastro.html", nil)
}
