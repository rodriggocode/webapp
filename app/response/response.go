package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Err struct {
	Erro string `json:"erro"`
}

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

func StatusCodeErr(w http.ResponseWriter, req *http.Response) {
	var erro Err
	json.NewDecoder(req.Body).Decode(&erro)
	JSON(w, req.StatusCode, erro)
}
