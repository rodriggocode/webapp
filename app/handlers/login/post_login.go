package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	resposta "webapp/app/response"
)

func PostLogin(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    req.FormValue("email"),
		"password": req.FormValue("password"),
	})
	if err != nil {
		resposta.JSON(w, http.StatusBadRequest, resposta.Err{Erro: err.Error()})
		return
	}

	urlApit := "https://devbook-zqaw.onrender.com/login"

	response, err := http.Post(urlApit, "application/json", bytes.NewBuffer(user))
	if err != nil {
		resposta.JSON(w, http.StatusInternalServerError, resposta.Err{Erro: err.Error()})
		return
	}

	token, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode, string(token))
}
