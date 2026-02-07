package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	respostas "webapp/app/response"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	users, err := json.Marshal(map[string]string{
		"user_name": req.FormValue("user_name"),
		"nick":      req.FormValue("nick"),
		"email":     req.FormValue("email"),
		"password":  req.FormValue("password"),
	})

	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.Err{Erro: err.Error()})
		return
	}

	response, err := http.Post("https://devbook-zqaw.onrender.com/create/user", "application/json", bytes.NewBuffer(users))
	if err != nil {

		respostas.JSON(w, http.StatusInternalServerError, respostas.Err{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.StatusCodeErr(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}
