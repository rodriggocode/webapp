package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/app/config"
	"webapp/app/cookies"
	"webapp/app/router"
	"webapp/app/utils"
)

func main() {

	config.LoadConfig()
	cookies.SetCookie()

	port := config.Port

	address := "http://localhost:7070"

	utils.LoadingTemplates()

	r := router.Gerar()

	fmt.Printf("Rodando WebApp %s\n", address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
