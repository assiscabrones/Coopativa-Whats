# Bot Nexum - Docker Deployment

Este guia explica como executar o Bot Nexum usando Docker e Docker Compose.

## 🐳 Pré-requisitos

- [Docker](https://docs.docker.com/get-docker/) (versão 20.10 ou superior)
- [Docker Compose](https://docs.docker.com/compose/install/) (versão 2.0 ou superior)

## 🚀 Instalação Rápida

### 1. Clone o Repositório
```bash
git clone <seu-repositorio>
cd gobot
```

### 2. Configure as Variáveis de Ambiente
```bash
# Copie o arquivo de exemplo
cp docker.env.example .env

# Edite o arquivo .env com suas configurações
nano .env
```

### 3. Execute o Script de Build
```bash
# Construir e iniciar o bot
./build-docker.sh build
./build-docker.sh run
```

## 📋 Configuração

### Variáveis de Ambiente

Edite o arquivo `.env` com as seguintes configurações:

```env
# Número do telefone para pairing (opcional)
PAIRING_NUMBER=5511999999999

# Lista de IDs dos owners (separados por vírgula)
OWNER=5511999999999,5511888888888

# Prefixo para comandos (não usado mais, mas mantido para compatibilidade)
PREFIX=!

# Se o bot é público (não usado mais, mas mantido para compatibilidade)
PUBLIC=true

# Configurações específicas do Docker
DATA_DIR=/app/data
SESSION_DIR=/app/session
```

## 🛠️ Comandos Disponíveis

### Script de Build (`build-docker.sh`)

```bash
# Construir a imagem Docker
./build-docker.sh build

# Iniciar o bot
./build-docker.sh run

# Ver status dos containers
./build-docker.sh status

# Ver logs em tempo real
./build-docker.sh logs

# Parar o bot
./build-docker.sh stop

# Limpar todos os recursos Docker
./build-docker.sh cleanup

# Mostrar ajuda
./build-docker.sh help
```

### Docker Compose

```bash
# Iniciar o bot
docker-compose up -d

# Ver logs
docker-compose logs -f

# Parar o bot
docker-compose down

# Reconstruir e iniciar
docker-compose up -d --build
```

### Docker Run (Alternativo)

```bash
# Construir a imagem
docker build -t bot-nexum:latest .

# Executar o container
docker run -d \
  --name bot-nexum \
  --restart unless-stopped \
  --env-file .env \
  -v "$(pwd)/data:/app/data" \
  -v "$(pwd)/session:/app/session" \
  bot-nexum:latest

# Ver logs
docker logs -f bot-nexum

# Parar o container
docker stop bot-nexum
docker rm bot-nexum
```

## 📁 Estrutura de Volumes

O Docker monta os seguintes volumes para persistência de dados:

```
./data/          → /app/data      # Banco de dados dos stages
./session/       → /app/session   # Sessão do WhatsApp
```

### Criando Diretórios
```bash
mkdir -p data session
```

## 🔧 Desenvolvimento

### Build Local
```bash
# Construir apenas a imagem
docker build -t bot-nexum:latest .

# Executar em modo interativo para debug
docker run -it --rm \
  --env-file .env \
  -v "$(pwd)/data:/app/data" \
  -v "$(pwd)/session:/app/session" \
  bot-nexum:latest
```

### Debug
```bash
# Entrar no container em execução
docker exec -it bot-nexum sh

# Ver logs detalhados
docker logs -f bot-nexum

# Verificar status do container
docker inspect bot-nexum
```

## 🚨 Troubleshooting

### Problemas Comuns

#### 1. Container não inicia
```bash
# Verificar logs
docker logs bot-nexum

# Verificar se as variáveis de ambiente estão corretas
docker exec -it bot-nexum env
```

#### 2. Problemas de permissão
```bash
# Verificar permissões dos volumes
ls -la data/ session/

# Corrigir permissões se necessário
sudo chown -R $USER:$USER data/ session/
```

#### 3. Banco de dados não persiste
```bash
# Verificar se o volume está montado
docker inspect bot-nexum | grep -A 10 "Mounts"

# Verificar se o diretório existe
docker exec -it bot-nexum ls -la /app/data/
```

#### 4. Sessão do WhatsApp perdida
```bash
# Verificar se o arquivo de sessão existe
docker exec -it bot-nexum ls -la /app/session/

# Recriar QR Code se necessário
docker restart bot-nexum
```

### Logs Úteis

```bash
# Logs em tempo real
docker-compose logs -f

# Logs das últimas 100 linhas
docker-compose logs --tail=100

# Logs de um serviço específico
docker-compose logs bot-nexum
```

## 🔒 Segurança

### Boas Práticas

1. **Não commite o arquivo `.env`** - Ele contém informações sensíveis
2. **Use usuário não-root** - O container roda como usuário `appuser`
3. **Volumes isolados** - Dados são persistidos em volumes específicos
4. **Restart policy** - Container reinicia automaticamente se falhar

### Variáveis Sensíveis

- `OWNER`: IDs dos administradores
- `PAIRING_NUMBER`: Número do telefone (se usado)

## 📊 Monitoramento

### Health Check
O container inclui um health check que verifica se o processo está rodando:

```bash
# Verificar status de saúde
docker inspect bot-nexum | grep -A 5 "Health"
```

### Métricas
```bash
# Uso de recursos
docker stats bot-nexum

# Informações do container
docker inspect bot-nexum
```

## 🔄 Atualizações

### Atualizar o Bot
```bash
# Parar o bot
./build-docker.sh stop

# Fazer pull das mudanças
git pull

# Reconstruir e iniciar
./build-docker.sh build
./build-docker.sh run
```

### Backup dos Dados
```bash
# Fazer backup dos dados
tar -czf backup-$(date +%Y%m%d).tar.gz data/ session/

# Restaurar backup
tar -xzf backup-20240101.tar.gz
```

## 🌐 Deploy em Produção

### Docker Swarm
```bash
# Inicializar swarm
docker swarm init

# Deploy do stack
docker stack deploy -c docker-compose.yml bot-nexum
```

### Kubernetes
```yaml
# Exemplo de deployment Kubernetes
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bot-nexum
spec:
  replicas: 1
  selector:
    matchLabels:
      app: bot-nexum
  template:
    metadata:
      labels:
        app: bot-nexum
    spec:
      containers:
      - name: bot-nexum
        image: bot-nexum:latest
        env:
        - name: OWNER
          value: "5511999999999"
        volumeMounts:
        - name: data
          mountPath: /app/data
        - name: session
          mountPath: /app/session
      volumes:
      - name: data
        persistentVolumeClaim:
          claimName: bot-nexum-data
      - name: session
        persistentVolumeClaim:
          claimName: bot-nexum-session
```

## 📞 Suporte

Para problemas específicos do Docker:

1. Verifique os logs: `./build-docker.sh logs`
2. Consulte a documentação do Docker
3. Abra uma issue no repositório

## 🎉 Conclusão

O Bot Nexum agora está totalmente compatível com Docker, oferecendo:

- ✅ Deploy simplificado
- ✅ Persistência de dados
- ✅ Restart automático
- ✅ Isolamento de ambiente
- ✅ Escalabilidade
- ✅ Monitoramento integrado

Execute `./build-docker.sh help` para ver todos os comandos disponíveis!
