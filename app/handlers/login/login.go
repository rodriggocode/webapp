package login

import (
	"net/http"
	"webapp/app/utils"
)

func Login(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "login.html", nil)
}
