FROM golang:1.25.7-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./main.go

FROM alpine/git

WORKDIR /app

COPY --from=builder /app/main .

COPY assets ./assets
COPY views ./views

EXPOSE 8080

CMD ["./main"]
