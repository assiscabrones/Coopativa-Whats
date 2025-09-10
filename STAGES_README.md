# Sistema de Stages - Bot Nexum

## Visão Geral

O sistema de stages substitui o antigo sistema de comandos, oferecendo uma experiência mais fluida e contextual para os usuários. Cada usuário tem um "stage" (estágio) atual que determina como o bot responde às suas mensagens.

## Características Principais

### 🎭 **Sistema de Stages**
- Cada usuário tem um stage atual salvo no banco de dados
- Navegação entre stages através de números ou palavras-chave
- Persistência automática do progresso do usuário
- Controle de permissões por stage (owner, grupo, privado)

### 💾 **Persistência de Dados**
- Banco de dados SQLite para armazenar stages dos usuários
- Dados específicos do usuário em cada stage
- Timestamps de criação e atualização

### 🔐 **Controle de Permissões**
- Stages exclusivos para owners
- Controle de acesso por tipo de chat (grupo/privado)
- Verificação automática de permissões

## Estrutura de Arquivos

```
src/
├── libs/
│   ├── stages.go          # Gerenciador principal do sistema
│   ├── types.go           # Estruturas de dados
│   └── message.go         # Serialização de mensagens
├── stages/
│   ├── index.go           # Importação de todos os stages
│   ├── basic/             # Stages básicos
│   │   ├── welcome.go     # Stage inicial
│   │   ├── menu.go        # Menu principal
│   │   ├── profile.go     # Gerenciamento de perfil
│   │   └── help.go        # Sistema de ajuda
│   └── admin/             # Stages administrativos
│       └── admin.go       # Painel de administração
└── handlers/
    └── message.go         # Handler de mensagens
```

## Como Funciona

### 1. **Inicialização**
- O sistema é inicializado no `hisoka.go`
- Cria tabela SQLite para armazenar stages dos usuários
- Registra todos os stages disponíveis

### 2. **Processamento de Mensagens**
- Cada mensagem é processada no contexto do stage atual do usuário
- O handler do stage determina como responder
- Navegação entre stages é gerenciada automaticamente

### 3. **Navegação**
- Usuários podem navegar digitando números (1, 2, 3) ou palavras-chave
- Cada stage define seus próprios "próximos stages" permitidos
- Validação automática de navegação

## Stages Disponíveis

### 🏠 **Welcome** (`welcome`)
- Stage inicial para todos os usuários
- Apresenta opções principais
- Navegação para: menu, config, help

### 📋 **Menu** (`menu`)
- Menu principal com funcionalidades
- Opção de administração para owners
- Navegação para: welcome, profile, games, tools, admin

### 👤 **Profile** (`profile`)
- Visualização e edição de perfil
- Informações do usuário
- Navegação para: menu, edit_profile

### ❓ **Help** (`help`)
- Sistema de ajuda e informações
- Explicações sobre navegação e stages
- Navegação para: welcome

### ⚙️ **Admin** (`admin`) - *Apenas Owners*
- Painel de administração
- Estatísticas do sistema
- Listagem de stages
- Navegação para: menu

## Criando Novos Stages

### 1. **Estrutura Básica**
```go
package basic

import (
    "hisoka/src/libs"
    // outros imports...
)

func meuStageHandler(conn *libs.IClient, m *libs.IMessage, userStage *libs.UserStage) bool {
    text := strings.ToLower(strings.TrimSpace(m.Text))
    
    switch text {
    case "1", "opcao1":
        // Lógica da opção 1
        return true
    case "0", "voltar":
        // Navegar para outro stage
        libs.ChangeUserStage(m.Sender.ToNonAD().User, "outro_stage")
        return true
    default:
        // Mostrar menu do stage
        m.Reply("Menu do stage...")
        return true
    }
}

func init() {
    libs.RegisterStage(&libs.Stage{
        ID:          "meu_stage",
        Name:        "Meu Stage",
        Description: "Descrição do stage",
        Handler:     meuStageHandler,
        NextStages:  []string{"outro_stage", "mais_um_stage"},
        IsOwner:     false,
        IsGroup:     false,
        IsPrivate:   false,
    })
}
```

### 2. **Registrar o Stage**
- Adicione o import no arquivo `stages/index.go`
- O stage será registrado automaticamente na inicialização

## Banco de Dados

### Tabela `user_stages`
```sql
CREATE TABLE user_stages (
    user_id TEXT PRIMARY KEY,
    current_stage TEXT NOT NULL,
    data TEXT,                    -- JSON com dados específicos do usuário
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL
);
```

## Variáveis de Ambiente

- `OWNER`: Lista de IDs de usuários owners (separados por vírgula)
- `PUBLIC`: Se o bot é público (não usado mais, mas mantido para compatibilidade)

## Migração do Sistema Antigo

O sistema de comandos foi completamente removido e substituído pelo sistema de stages:

- ❌ `src/commands/` - Removido
- ❌ `src/libs/commands.go` - Removido
- ✅ `src/stages/` - Novo sistema
- ✅ `src/libs/stages.go` - Gerenciador de stages

## Vantagens do Sistema de Stages

1. **Contexto Persistente**: Cada usuário mantém seu estado
2. **Navegação Intuitiva**: Interface mais amigável
3. **Flexibilidade**: Fácil adição de novos stages
4. **Controle de Permissões**: Controle granular de acesso
5. **Persistência**: Dados salvos automaticamente
6. **Escalabilidade**: Sistema modular e extensível

## Exemplo de Uso

1. Usuário envia mensagem → Bot verifica stage atual
2. Handler do stage processa a mensagem
3. Usuário navega digitando opções
4. Stage é atualizado e salvo no banco
5. Próxima mensagem usa o novo stage

O sistema oferece uma experiência muito mais rica e contextual para os usuários, mantendo o estado da conversa e permitindo fluxos complexos de interação.
