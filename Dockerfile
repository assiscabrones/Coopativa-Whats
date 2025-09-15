# Use a imagem oficial do Go como base
FROM golang:1.21-alpine AS builder

# Instalar dependências necessárias para compilação
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=1 GOOS=linux go build -trimpath -buildvcs=false -ldflags="-s -w" -o main .

# Imagem final
FROM alpine:3.20

# Instalar dependências de runtime
RUN apk --no-cache add ca-certificates sqlite-libs tzdata

# Criar usuário não-root para segurança
RUN adduser -D -s /bin/sh appuser

# Definir diretório de trabalho
WORKDIR /app

# Copiar o binário compilado
COPY --from=builder /app/main .

# (Opcional) Copiar arquivos de configuração de exemplo
# COPY --from=builder /app/config.example .

# Criar diretórios para dados persistentes
RUN mkdir -p /app/data && chown -R appuser:appuser /app

# Mudar para usuário não-root
USER appuser

# Expor porta (se necessário para health checks)
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]
