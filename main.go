package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/app/router"
	"webapp/app/utils"
)

func main() {

	utils.LoadingTemplates()

	r := router.Gerar()

	fmt.Println("Rodando WebApp!")
	log.Fatal(http.ListenAndServe(":7000", r))
}
