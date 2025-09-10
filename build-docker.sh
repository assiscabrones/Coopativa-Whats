#!/bin/bash

# Script para build e deploy do Bot Nexum com Docker

set -e

echo "🐳 Bot Nexum - Docker Build Script"
echo "=================================="

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Função para imprimir mensagens coloridas
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Verificar se Docker está instalado
if ! command -v docker &> /dev/null; then
    print_error "Docker não está instalado. Por favor, instale o Docker primeiro."
    exit 1
fi

# Verificar se Docker Compose está instalado
if ! command -v docker-compose &> /dev/null; then
    print_warning "Docker Compose não está instalado. Usando 'docker compose'..."
    COMPOSE_CMD="docker compose"
else
    COMPOSE_CMD="docker-compose"
fi

# Função para criar diretórios necessários
create_directories() {
    print_status "Criando diretórios necessários..."
    mkdir -p data session
    print_success "Diretórios criados com sucesso"
}

# Função para verificar arquivo de configuração
check_config() {
    if [ ! -f ".env" ]; then
        print_warning "Arquivo .env não encontrado. Copiando exemplo..."
        if [ -f "docker.env.example" ]; then
            cp docker.env.example .env
            print_success "Arquivo .env criado a partir do exemplo"
            print_warning "Por favor, edite o arquivo .env com suas configurações antes de continuar"
            return 1
        else
            print_error "Arquivo docker.env.example não encontrado"
            return 1
        fi
    fi
    return 0
}

# Função para build da imagem
build_image() {
    print_status "Construindo imagem Docker..."
    docker build -t bot-nexum:latest .
    print_success "Imagem construída com sucesso"
}

# Função para executar com Docker Compose
run_with_compose() {
    print_status "Iniciando Bot Nexum com Docker Compose..."
    $COMPOSE_CMD up -d
    print_success "Bot Nexum iniciado com sucesso"
    print_status "Use '$COMPOSE_CMD logs -f' para ver os logs"
    print_status "Use '$COMPOSE_CMD down' para parar o bot"
}

# Função para executar com Docker run
run_with_docker() {
    print_status "Iniciando Bot Nexum com Docker run..."
    docker run -d \
        --name bot-nexum \
        --restart unless-stopped \
        --env-file .env \
        -v "$(pwd)/data:/app/data" \
        -v "$(pwd)/session:/app/session" \
        bot-nexum:latest
    print_success "Bot Nexum iniciado com sucesso"
    print_status "Use 'docker logs -f bot-nexum' para ver os logs"
    print_status "Use 'docker stop bot-nexum' para parar o bot"
}

# Função para mostrar status
show_status() {
    print_status "Status dos containers:"
    if command -v docker-compose &> /dev/null; then
        $COMPOSE_CMD ps
    else
        docker ps --filter "name=bot-nexum"
    fi
}

# Função para mostrar logs
show_logs() {
    print_status "Mostrando logs do Bot Nexum:"
    if command -v docker-compose &> /dev/null; then
        $COMPOSE_CMD logs -f
    else
        docker logs -f bot-nexum
    fi
}

# Função para parar o bot
stop_bot() {
    print_status "Parando Bot Nexum..."
    if command -v docker-compose &> /dev/null; then
        $COMPOSE_CMD down
    else
        docker stop bot-nexum
        docker rm bot-nexum
    fi
    print_success "Bot Nexum parado com sucesso"
}

# Função para limpeza
cleanup() {
    print_status "Limpando recursos Docker..."
    if command -v docker-compose &> /dev/null; then
        $COMPOSE_CMD down --rmi all --volumes --remove-orphans
    else
        docker stop bot-nexum 2>/dev/null || true
        docker rm bot-nexum 2>/dev/null || true
        docker rmi bot-nexum:latest 2>/dev/null || true
    fi
    print_success "Limpeza concluída"
}

# Menu principal
case "${1:-build}" in
    "build")
        create_directories
        if check_config; then
            build_image
            print_success "Build concluído! Use './build-docker.sh run' para iniciar o bot"
        fi
        ;;
    "run")
        if [ ! -f ".env" ]; then
            print_error "Arquivo .env não encontrado. Execute './build-docker.sh build' primeiro"
            exit 1
        fi
        if [ -f "docker-compose.yml" ]; then
            run_with_compose
        else
            run_with_docker
        fi
        ;;
    "status")
        show_status
        ;;
    "logs")
        show_logs
        ;;
    "stop")
        stop_bot
        ;;
    "cleanup")
        cleanup
        ;;
    "help"|"-h"|"--help")
        echo "Uso: $0 [comando]"
        echo ""
        echo "Comandos disponíveis:"
        echo "  build    - Constrói a imagem Docker (padrão)"
        echo "  run      - Inicia o bot com Docker"
        echo "  status   - Mostra o status dos containers"
        echo "  logs     - Mostra os logs do bot"
        echo "  stop     - Para o bot"
        echo "  cleanup  - Remove todos os recursos Docker"
        echo "  help     - Mostra esta ajuda"
        ;;
    *)
        print_error "Comando desconhecido: $1"
        echo "Use '$0 help' para ver os comandos disponíveis"
        exit 1
        ;;
esac
