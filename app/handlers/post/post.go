package post

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/app/request"
	respostas "webapp/app/response"
)

func CreatePost(w http.ResponseWriter, req *http.Request) {
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.Err{Erro: err.Error()})
		return
	}

	post, err := json.Marshal(map[string]string{
		"title":   body.Title,
		"content": body.Content,
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
