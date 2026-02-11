package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL   = ""
	Port     = 0
	HasKey   []byte // autenticar o cookie
	BlockKey []byte // criptografar os dados do cookie
)

func LoadConfig() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal()
	}
	APIURL = os.Getenv("API_URL")
	HasKey = []byte(os.Getenv("HAS_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
