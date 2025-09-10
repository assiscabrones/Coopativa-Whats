## Langkah Instalasi

1. [`Git`](https://git-scm.com/downloads)
2. [`Go`](https://go.dev/doc/install)
3. Gcc for build

> Gcc for windows on [`Here`](https://dev.to/gamegods3/how-to-install-gcc-in-windows-10-the-easier-way-422j)

> Gcc for linux (Ubuntu) on [`Here`](https://linuxize.com/post/how-to-install-gcc-on-ubuntu-20-04/)

## Clone Repository 
```sh
git clone https://github.com/DikaArdnt/go-readsw
```

## go to the folder 
```sh
cd go-readsw
```

## Install Dependencies
```sh
go get all
```

## Build (Optional)
```sh
go build .

# Run
./hisoka.exe # for Windows
hisoka.exe # for linux
```

## Run
```sh
go run main.go
```

## 🐳 Docker Deployment

O Bot Nexum agora suporta deployment com Docker para facilitar a instalação e execução.

### Pré-requisitos para Docker
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Instalação Rápida com Docker
```bash
# Clone o repositório
git clone <seu-repositorio>
cd gobot

# Configure as variáveis de ambiente
cp env.example .env
# Edite o arquivo .env com suas configurações

# Execute o script de build
./build-docker.sh build
./build-docker.sh run
```

### Comandos Docker Disponíveis
```bash
# Construir a imagem
./build-docker.sh build

# Iniciar o bot
./build-docker.sh run

# Ver logs
./build-docker.sh logs

# Parar o bot
./build-docker.sh stop

# Ver status
./build-docker.sh status
```

### Docker Compose
```bash
# Iniciar com Docker Compose
docker-compose up -d

# Ver logs
docker-compose logs -f

# Parar
docker-compose down
```

Para mais informações sobre Docker, consulte [DOCKER_README.md](DOCKER_README.md).

## Thanks To
- [tulir](https://github.com/tulir)
- [vnia](https://github.com/fckvania)# Coopativa-Whats
