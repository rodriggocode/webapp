package post

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/app/request"
	respostas "webapp/app/response"
)

func CreatePost(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	post, err := json.Marshal(map[string]string{
		"title":   req.FormValue("title"),
		"content": req.FormValue("content"),
	})

	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.Err{Erro: err.Error()})
		return
	}

	response, err := request.RequestAuth(req, http.MethodPost, "https://devbook-zqaw.onrender.com/publicacao/criar", bytes.NewBuffer(post))

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
