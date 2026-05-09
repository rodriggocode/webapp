// aqui vai ser so aparte que vai renderizar os htmls da views
package pages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/app/cookies"
	"webapp/app/models"
	"webapp/app/request"
	"webapp/app/response"
	utils "webapp/app/utils"

	"github.com/golang-jwt/jwt/v5"
)

func LoadPageCreateUser(w http.ResponseWriter, req *http.Request) {
	utils.ExecuterTemplate(w, "cadastro.html", nil)
}

func LoadHomePage(w http.ResponseWriter, req *http.Request) {
	url := "https://devbook-zqaw.onrender.com/publicacoes"
	resp, err := request.RequestAuth(req, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Err{Erro: err.Error()})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		response.StatusCodeErr(w, resp)
		return
	}

	var posts []models.Posts
	if err = json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.Err{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.ReadToken(req)

	token, _, err := new(jwt.Parser).ParseUnverified(cookie, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)
	userId := uint64(claims["id_user"].(float64))
	utils.ExecuterTemplate(w, "home.html", struct {
		Posts  []models.Posts
		UserID uint64
	}{
		Posts:  posts,
		UserID: userId,
	})
}
