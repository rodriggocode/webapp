package home

import (
	"net/http"
	"webapp/app/utils"
)

func HomePage(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "home.tml", nil)
}
