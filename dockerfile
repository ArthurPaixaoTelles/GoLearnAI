# ============================
# Stage 1: Build
# ============================
FROM golang:1.22-alpine AS builder

LABEL stage="builder"

# Instala dependências básicas
RUN apk add --no-cache git ca-certificates

# Diretório de trabalho
WORKDIR /app

# Copia mod files e baixa dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila binário otimizado
RUN go build -ldflags="-s -w" -o app main.go


# ============================
# Stage 2: Runtime
# ============================
FROM alpine:3.20

# Adiciona certificados SSL e dependências mínimas
RUN apk add --no-cache ca-certificates tzdata \
    && adduser -D appuser

# Diretório de trabalho
WORKDIR /home/appuser

# Copia o binário do estágio anterior
COPY --from=builder /app/app .

# Define permissões e muda o usuário
RUN chown appuser:appuser ./app
USER appuser

# Define variáveis de ambiente padrão
ENV PORT=8080
ENV HF_API_KEY=""

# Expondo a porta da API
EXPOSE 8080

# Labels úteis para rastreamento
LABEL org.opencontainers.image.title="Go LLM API" \
      org.opencontainers.image.description="API em Go com integração HuggingFace" \
      org.opencontainers.image.version="1.0" \
      maintainer="seu_nome <seu_email>"

# Comando para iniciar o servidor
CMD ["./app"]
