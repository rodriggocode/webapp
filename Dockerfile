# Dockerfile mínimo para Fly.io
ARG GO_VERSION=1.25
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /run-app ./main.go

FROM debian:bookworm
COPY --from=builder /run-app /usr/local/bin/
EXPOSE 7070
CMD ["run-app"]

