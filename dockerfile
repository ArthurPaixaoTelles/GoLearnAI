# Stage 1: Builder
FROM golang:1.22-alpine AS builder

# Dependências básicas
RUN apk add --no-cache git ca-certificates

# Diretório de trabalho
WORKDIR /app

# Copia go.mod e go.sum e baixa dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código
COPY . .

# Compila o binário
RUN go build -o app main.go

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /root/

# Copia binário do builder
COPY --from=builder /app/app .

# Porta do container
EXPOSE 8080

# Variável de ambiente do HuggingFace será passada na execução
# CMD inicializa o servidor
CMD ["./app"]
