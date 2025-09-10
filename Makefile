# Bot Nexum - Makefile

.PHONY: help build run test clean docker-build docker-run docker-stop docker-logs

# Variáveis
BINARY_NAME=bot
DOCKER_IMAGE=bot-nexum
DOCKER_TAG=latest

# Cores para output
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color

help: ## Mostra esta ajuda
	@echo "$(GREEN)Bot Nexum - Comandos Disponíveis$(NC)"
	@echo "=================================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'

build: ## Compila o projeto
	@echo "$(GREEN)Compilando Bot Nexum...$(NC)"
	@go build -o $(BINARY_NAME) .
	@echo "$(GREEN)Compilação concluída!$(NC)"

run: ## Executa o bot
	@echo "$(GREEN)Executando Bot Nexum...$(NC)"
	@go run main.go

test: ## Executa os testes
	@echo "$(GREEN)Executando testes...$(NC)"
	@go test ./...

clean: ## Limpa arquivos de build
	@echo "$(YELLOW)Limpando arquivos de build...$(NC)"
	@rm -f $(BINARY_NAME)
	@rm -f *.db
	@rm -f *.log
	@echo "$(GREEN)Limpeza concluída!$(NC)"

# Comandos Docker
docker-build: ## Constrói a imagem Docker
	@echo "$(GREEN)Construindo imagem Docker...$(NC)"
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "$(GREEN)Imagem construída com sucesso!$(NC)"

docker-run: ## Executa o bot com Docker
	@echo "$(GREEN)Executando Bot Nexum com Docker...$(NC)"
	@docker run -d \
		--name $(DOCKER_IMAGE) \
		--restart unless-stopped \
		--env-file .env \
		-v "$(PWD)/data:/app/data" \
		-v "$(PWD)/session:/app/session" \
		$(DOCKER_IMAGE):$(DOCKER_TAG)
	@echo "$(GREEN)Bot Nexum iniciado com Docker!$(NC)"

docker-stop: ## Para o container Docker
	@echo "$(YELLOW)Parando Bot Nexum...$(NC)"
	@docker stop $(DOCKER_IMAGE) || true
	@docker rm $(DOCKER_IMAGE) || true
	@echo "$(GREEN)Bot Nexum parado!$(NC)"

docker-logs: ## Mostra os logs do container Docker
	@echo "$(GREEN)Mostrando logs do Bot Nexum...$(NC)"
	@docker logs -f $(DOCKER_IMAGE)

docker-compose-up: ## Inicia com Docker Compose
	@echo "$(GREEN)Iniciando com Docker Compose...$(NC)"
	@docker-compose up -d
	@echo "$(GREEN)Bot Nexum iniciado com Docker Compose!$(NC)"

docker-compose-down: ## Para com Docker Compose
	@echo "$(YELLOW)Parando com Docker Compose...$(NC)"
	@docker-compose down
	@echo "$(GREEN)Bot Nexum parado!$(NC)"

docker-compose-logs: ## Mostra logs do Docker Compose
	@echo "$(GREEN)Mostrando logs do Docker Compose...$(NC)"
	@docker-compose logs -f

# Comandos de desenvolvimento
dev: ## Executa em modo de desenvolvimento
	@echo "$(GREEN)Executando em modo de desenvolvimento...$(NC)"
	@go run main.go

install: ## Instala dependências
	@echo "$(GREEN)Instalando dependências...$(NC)"
	@go mod tidy
	@go mod download
	@echo "$(GREEN)Dependências instaladas!$(NC)"

fmt: ## Formata o código
	@echo "$(GREEN)Formatando código...$(NC)"
	@go fmt ./...
	@echo "$(GREEN)Código formatado!$(NC)"

lint: ## Executa o linter
	@echo "$(GREEN)Executando linter...$(NC)"
	@golangci-lint run || echo "$(YELLOW)Linter não encontrado. Instale com: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(NC)"

# Comandos de setup
setup: ## Configura o ambiente de desenvolvimento
	@echo "$(GREEN)Configurando ambiente de desenvolvimento...$(NC)"
	@mkdir -p data session
	@cp env.example .env || echo "$(YELLOW)Arquivo env.example não encontrado$(NC)"
	@echo "$(GREEN)Ambiente configurado!$(NC)"
	@echo "$(YELLOW)Edite o arquivo .env com suas configurações$(NC)"

# Comandos de backup
backup: ## Faz backup dos dados
	@echo "$(GREEN)Fazendo backup dos dados...$(NC)"
	@tar -czf backup-$(shell date +%Y%m%d-%H%M%S).tar.gz data/ session/ *.db 2>/dev/null || echo "$(YELLOW)Nenhum dado para backup$(NC)"
	@echo "$(GREEN)Backup concluído!$(NC)"

# Comandos de status
status: ## Mostra o status do bot
	@echo "$(GREEN)Status do Bot Nexum:$(NC)"
	@echo "========================"
	@if [ -f $(BINARY_NAME) ]; then echo "$(GREEN)✓ Binário compilado$(NC)"; else echo "$(RED)✗ Binário não encontrado$(NC)"; fi
	@if [ -f .env ]; then echo "$(GREEN)✓ Arquivo .env encontrado$(NC)"; else echo "$(RED)✗ Arquivo .env não encontrado$(NC)"; fi
	@if [ -d data ]; then echo "$(GREEN)✓ Diretório data existe$(NC)"; else echo "$(YELLOW)⚠ Diretório data não existe$(NC)"; fi
	@if [ -d session ]; then echo "$(GREEN)✓ Diretório session existe$(NC)"; else echo "$(YELLOW)⚠ Diretório session não existe$(NC)"; fi
	@if docker ps --format "table {{.Names}}" | grep -q $(DOCKER_IMAGE); then echo "$(GREEN)✓ Container Docker rodando$(NC)"; else echo "$(RED)✗ Container Docker não está rodando$(NC)"; fi

# Comando padrão
.DEFAULT_GOAL := help
