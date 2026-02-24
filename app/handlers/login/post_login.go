package login

import (
	"bytes"
	"encoding/json"

	"net/http"
	"webapp/app/cookies"
	"webapp/app/models"
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

	urlApi := "https://devbook-zqaw.onrender.com/login"

	response, err := http.Post(urlApi, "application/json", bytes.NewBuffer(user))
	if err != nil {
		resposta.JSON(w, http.StatusInternalServerError, resposta.Err{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		resposta.StatusCodeErr(w, response)
		return
	}

	var dateAuth models.DadoAtuh
	if err = json.NewDecoder(response.Body).Decode(&dateAuth); err != nil {
		resposta.JSON(w, http.StatusUnprocessableEntity, resposta.Err{Erro: err.Error()})
		return
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		cookies.SetCookie(w, dateAuth.Token)
		w.WriteHeader(http.StatusOK)

		resposta.JSON(w, http.StatusOK, map[string]string{
			"redirect": "/home",
		})

		return
	}

	resposta.JSON(w, http.StatusOK, nil)

}
