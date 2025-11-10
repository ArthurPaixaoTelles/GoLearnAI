# Stage 1: Builder
FROM golang:1.22-alpine AS builder

# Instala dependências básicas
RUN apk add --no-cache git ca-certificates

# Define diretório de trabalho
WORKDIR /app

# Copia arquivos go.mod/go.sum e baixa dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código
COPY . .

# Compila o binário
RUN go build -o app main.go

# Stage 2: Runtime
FROM alpine:latest

# Diretório de trabalho
WORKDIR /root/

# Copia binário do builder
COPY --from=builder /app/app .

# Copia .env se necessário (opcional)
# COPY --from=builder /app/.env .env

# Porta que o container vai expor
EXPOSE 8080

# Comando para rodar o servidor
CMD ["./app"]
