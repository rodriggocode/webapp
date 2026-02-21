package home

import (
	"net/http"
	"webapp/app/utils"
)

func Home(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "home.tml", nil)
}
