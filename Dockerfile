# Use a imagem oficial do Go para construir a aplicação
FROM golang:1.22-alpine AS builder

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia os arquivos go.mod e go.sum para resolver as dependências
COPY go.mod ./
COPY go.sum ./

# Baixa as dependências do Go (módulos)
# Isso ajuda a acelerar builds subsequentes, pois esta camada será cacheada
RUN go mod download

# Copia todo o código-fonte da sua aplicação
COPY . .

# Constrói a aplicação Go
# - CGO_ENABLED=0: Desabilita CGO para criar um binário estático, que é mais fácil de empacotar
# - -o main: Define o nome do executável como 'main'
# - -ldflags "-s -w": Reduz o tamanho do binário removendo informações de depuração
# - ./main.go: O arquivo principal da sua aplicação Go
RUN go build -o main ./main.go


# Use uma imagem base menor para a execução (runtime) da aplicação
FROM alpine/git

# Define o diretório de trabalho dentro do contêiner de execução
WORKDIR /app

# Copia o binário compilado da fase de construção
COPY --from=builder /app/main .

# Copia os arquivos estáticos e templates necessários (assumindo que sua aplicação os sirva)
# Ajuste esses caminhos conforme a estrutura real do seu projeto, se necessário.
# Pela sua estrutura, parece que 'app/assets' e 'views' são importantes.
COPY assets ./assets
COPY views ./views

# Adicione outros arquivos que sua aplicação precise em runtime, por exemplo:
# COPY .env ./.env # APENAS SE VOCÊ QUISER CARREGAR .env NO CONTAINER (geralmente não recomendado para produção)

# O Fly.io injeta a porta na variável de ambiente PORT.
# Certifique-se de que sua aplicação Go esteja lendo esta variável (os.Getenv("PORT")).
# Se sua aplicação espera "APP_PORT", você precisará configurar isso no fly.toml ou no código.
EXPOSE 8080 # Porta padrão para o Go, o Fly.io mapeia automaticamente

# Comando para rodar a aplicação quando o contêiner iniciar
CMD ["./main"]
