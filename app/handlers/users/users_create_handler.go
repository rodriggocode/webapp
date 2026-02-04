package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	users, err := json.Marshal(map[string]string{
		"user_name": req.FormValue("user_name"),
		"nick":      req.FormValue("nick"),
		"email":     req.FormValue("email"),
		"password":  req.FormValue("password"),
	})

	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("https://devbook-zqaw.onrender.com/create/user", "application/json", bytes.NewBuffer(users))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println(response.Body)

}
