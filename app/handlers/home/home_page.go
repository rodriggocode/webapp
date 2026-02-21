package home

import "net/http"

func HomePage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Pagina de Home"))
}
