package login

import (
	"bytes"
	"encoding/json"
	"log"

	"net/http"
	"webapp/app/cookies"
	"webapp/app/models"
	"webapp/app/response"
	resposta "webapp/app/response"
)

func PostLogin(w http.ResponseWriter, req *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(req.Body).Decode(&credentials); err != nil {
		resposta.JSON(w, http.StatusBadRequest, resposta.Err{Erro: err.Error()})
		return
	}

	log.Printf("Email recebido: %s", credentials.Email)

	user, err := json.Marshal(map[string]string{
		"email":    credentials.Email,
		"password": credentials.Password,
	})

	if err != nil {
		resposta.JSON(w, http.StatusBadRequest, resposta.Err{Erro: err.Error()})
		return
	}
	urlApi := "https://devbook-zqaw.onrender.com/login"
	resp, err := http.Post(urlApi, "application/json", bytes.NewBuffer(user))
	if err != nil {
		resposta.JSON(w, http.StatusInternalServerError, response.Err{Erro: err.Error()})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		resposta.StatusCodeErr(w, resp)
		return
	}

	var dateAuth models.DadoAtuh
	if err = json.NewDecoder(resp.Body).Decode(&dateAuth); err != nil {
		resposta.JSON(w, http.StatusUnprocessableEntity, resposta.Err{Erro: err.Error()})
		return
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		cookies.SetCookie(w, dateAuth.Token)
		resposta.JSON(w, http.StatusOK, map[string]string{
			"redirect": "/home",
		})
		return
	}

	resposta.JSON(w, http.StatusOK, nil)
}
