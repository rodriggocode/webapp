package cookies

import (
	"fmt"
	"net/http"
	"webapp/app/config"
)

func SetCookie(w http.ResponseWriter, req *http.Request) {
	cookie_env := config.HasKey
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    string(cookie_env), // aqui eu vou coloca o que vem do meu .env
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "cookie has been set!")
}
