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
		log.Println("Aviso: Nao foi possivel carregar o arquivo .env")
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "7070"
		log.Println("Variavel de ambiente PORT nao definida, usando padrao 7070")
	}

	Port, err = strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Erro ao converter PORT '%s' para numero, %v", portStr, err)
	}

	log.Printf("Aplicao configurada para escutar na port %d", Port)

	APIURL = os.Getenv("API_URL")
	HasKey = []byte(os.Getenv("HAS_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
