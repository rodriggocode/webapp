package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/app/config"
	"webapp/app/router"
	"webapp/app/utils"
)

func main() {

	config.LoadConfig()

	port := config.Port

	address := fmt.Sprintf("http://0.0.0.0:%d", port)

	utils.LoadingTemplates()

	r := router.Gerar()

	fmt.Printf("Rodando WebApp em %s\n", address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), r))
}
