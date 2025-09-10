package libs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var stages map[string]*Stage
var db *sql.DB

// Inicializa o sistema de stages
func InitStages() error {
	stages = make(map[string]*Stage)
	
	// Obter diretório de dados das variáveis de ambiente
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "."
	}
	
	// Criar diretório se não existir
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}
	
	dbPath := dataDir + "/stages.db"
	
	// Conecta ao banco de dados
	var err error
	db, err = sql.Open("sqlite3", "file:"+dbPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	
	// Cria a tabela de usuários se não existir
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS user_stages (
		user_id TEXT PRIMARY KEY,
		current_stage TEXT NOT NULL,
		data TEXT,
		created_at INTEGER NOT NULL,
		updated_at INTEGER NOT NULL
	);`
	
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	
	// Registra stages básicos se não foram registrados automaticamente
	registerBasicStages()
	
	return nil
}

// Registra stages básicos manualmente se necessário
func registerBasicStages() {
	// Registra o stage padrão
	RegisterStage(&Stage{
		ID:          "default",
		Name:        "Menu Principal",
		Description: "Menu principal de atendimento",
		Handler:     defaultHandler,
		NextStages:  []string{
			"adesao", "aplicativo", "capital", "emprestimos", 
			"parcerias", "consultoria", "excolaborador", 
			"negociacao", "informe", "duvidas",
		},
		IsOwner:     false,
		IsGroup:     false,
		IsPrivate:   false,
	})
	
	// Registra o stage de adesão
	RegisterStage(&Stage{
		ID:          "adesao",
		Name:        "Adesão",
		Description: "Processo de adesão à Ativa Grupo SBF",
		Handler:     adesaoHandler,
		NextStages:  []string{"default"},
		IsOwner:     false,
		IsGroup:     false,
		IsPrivate:   false,
	})
	
	// Registra o stage de aplicativo/senha
	RegisterStage(&Stage{
		ID:          "aplicativo",
		Name:        "Aplicativo ou Senha",
		Description: "Ajuda com aplicativo e senhas de acesso",
		Handler:     aplicativoHandler,
		NextStages:  []string{"default"},
		IsOwner:     false,
		IsGroup:     false,
		IsPrivate:   false,
	})
}

// Handler do stage default
func defaultHandler(conn *IClient, m *IMessage, userStage *UserStage) bool {
	// Teste simples primeiro
	fmt.Println("🚀 [DEFAULT] TESTE - Handler INICIADO!")
	fmt.Printf("🚀 [DEFAULT] Handler INICIADO para usuário %s\n", m.Sender.ToNonAD().User)
	fmt.Printf("🚀 [DEFAULT] Parâmetros: conn=%v, m=%v, userStage=%v\n", conn != nil, m != nil, userStage != nil)
	
	text := strings.ToLower(strings.TrimSpace(m.Text))
	fmt.Printf("🔍 [DEFAULT] Handler recebeu: '%s' do usuário %s\n", text, m.Sender.ToNonAD().User)
	
	switch text {
	case "1", "adesão", "adesao":
		fmt.Printf("🔄 [DEFAULT] Usuário quer ir para adesão\n")
		err := ChangeUserStage(m.Sender.ToNonAD().User, "adesao")
		if err != nil {
			fmt.Printf("❌ [DEFAULT] Erro ao mudar stage: %s\n", err.Error())
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		fmt.Printf("✅ [DEFAULT] Stage mudado para 'adesao'\n")
		adesaoStage := GetStage("adesao")
		if adesaoStage != nil && adesaoStage.Handler != nil {
			fmt.Printf("🔄 [DEFAULT] Executando handler do stage adesao\n")
			userStage, _ := GetUserStage(m.Sender.ToNonAD().User)
			adesaoStage.Handler(conn, m, userStage)
			fmt.Printf("✅ [DEFAULT] Handler do adesao executado\n")
		} else {
			fmt.Printf("❌ [DEFAULT] Stage adesao não encontrado ou sem handler\n")
		}
		return true
		
	case "2", "aplicativo", "senha", "acesso":
		// Navega para stage de aplicativo/senha e executa o handler imediatamente
		err := ChangeUserStage(m.Sender.ToNonAD().User, "aplicativo")
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		aplicativoStage := GetStage("aplicativo")
		if aplicativoStage != nil && aplicativoStage.Handler != nil {
			fmt.Printf("🔄 [DEFAULT] Executando handler do stage aplicativo\n")
			userStage, _ := GetUserStage(m.Sender.ToNonAD().User)
			aplicativoStage.Handler(conn, m, userStage)
			fmt.Printf("✅ [DEFAULT] Handler do aplicativo executado\n")
		}
		return true

	case "3", "capital", "investimento":
		// Navega para stage de capital
		err := ChangeUserStage(m.Sender.ToNonAD().User, "capital")
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "4", "empréstimos", "emprestimos":
		// Navega para stage de empréstimos
		err := ChangeUserStage(m.Sender.ToNonAD().User, "emprestimos")
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "5", "parcerias":
		// Navega para stage de parcerias
		err := ChangeUserStage(m.Sender.ToNonAD().User, "parcerias")
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "6", "consultoria", "financeira":
		// Navega para stage de consultoria
		err := ChangeUserStage(m.Sender.ToNonAD().User, "consultoria")
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "7", "ex-colaborador", "excolaborador":
		// Navega para stage de ex-colaborador
		err := ChangeUserStage(m.Sender.ToNonAD().User, "excolaborador")
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "8", "negociação", "negociacao", "dívidas", "dividas":
		// Navega para stage de negociação de dívidas
		err := ChangeUserStage(m.Sender.ToNonAD().User, "negociacao")
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "9", "informe", "rendimentos":
		// Navega para stage de informe de rendimentos
		err := ChangeUserStage(m.Sender.ToNonAD().User, "informe")
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "10", "dúvida", "duvida", "não encontrou", "nao encontrou":
		// Navega para stage de dúvidas
		err := ChangeUserStage(m.Sender.ToNonAD().User, "duvidas")
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "11", "encerrar", "sair", "fim":
		// Encerra o atendimento
		m.Reply(`👋 *Atendimento encerrado!*

Obrigado por entrar em contato conosco.

Se precisar de mais alguma coisa, é só me chamar novamente! 😊`)
		return true

	default:
		fmt.Printf("🔄 [DEFAULT] Enviando mensagem padrão do menu\n")
		// Mostra o menu principal
		message := fmt.Sprintf(`🏢 *Olá! Bem-vindo ao Whatsapp da Ativa Grupo SBF 😃*

Olá, %s! 👋
Informamos que as mensagens deste canal devem ser apenas de texto. Não atendemos mensagens de voz ou ligações.

Escolha a opção desejada para atendimento:

📋 *MENU PRINCIPAL*

1️⃣ *Adesão* - Informações sobre adesão
2️⃣ *Aplicativo ou Senha* - Acesso ao sistema
3️⃣ *Capital (Investimento)* - Produtos de investimento
4️⃣ *Empréstimos* - Soluções de crédito
5️⃣ *Parcerias* - Oportunidades de parceria
6️⃣ *Consultoria Financeira* - Orientação especializada
7️⃣ *Ex-colaborador* - Atendimento para ex-funcionários
8️⃣ *Negociação de Dívidas* - Ex-colaborador
9️⃣ *Informe de Rendimentos* - Documentos fiscais
🔟 *Não encontrou sua dúvida?* - Atendimento personalizado
1️⃣1️⃣ *Encerrar Atendimento* - Finalizar conversa

💡 *Como usar:*
• Digite o *número* da opção (ex: 1, 2, 3...)
• Digite o *nome* da opção (ex: adesão, empréstimos)
• Use palavras-chave como *sair* ou *encerrar*

Escolha uma opção para continuar! ⬇️`, m.Info.PushName)
		
		m.Reply(message)
		return true
	}
	
	fmt.Printf("⚠️ [DEFAULT] Nenhum caso foi executado para: '%s'\n", text)
	return false
}

// Handler do stage de adesão
func adesaoHandler(conn *IClient, m *IMessage, userStage *UserStage) bool {
	fmt.Printf("🚀 [ADESAO] Handler INICIADO para usuário %s\n", m.Sender.ToNonAD().User)
	fmt.Printf("🚀 [ADESAO] Parâmetros: conn=%v, m=%v, userStage=%v\n", conn != nil, m != nil, userStage != nil)
	
	text := strings.ToLower(strings.TrimSpace(m.Text))
	
	fmt.Printf("🔍 [ADESAO] Handler recebeu: '%s' do usuário %s\n", text, m.Sender.ToNonAD().User)
	fmt.Printf("🔍 [ADESAO] Texto processado: '%s'\n", text)
	
	switch text {
	case "0", "voltar", "menu", "início", "inicio":
		fmt.Printf("🔄 [ADESAO] Usuário quer voltar ao menu principal\n")
		err := ChangeUserStage(m.Sender.ToNonAD().User, "default")
		if err != nil {
			fmt.Printf("❌ [ADESAO] Erro ao mudar stage: %s\n", err.Error())
			m.Reply("❌ Erro ao voltar: " + err.Error())
			return false
		}
		fmt.Printf("✅ [ADESAO] Stage mudado para 'default'\n")
		defaultStage := GetStage("default")
		if defaultStage != nil && defaultStage.Handler != nil {
			fmt.Printf("🔄 [ADESAO] Executando handler do stage default\n")
			userStage, _ := GetUserStage(m.Sender.ToNonAD().User)
			defaultStage.Handler(conn, m, userStage)
			fmt.Printf("✅ [ADESAO] Handler do default executado\n")
		} else {
			fmt.Printf("❌ [ADESAO] Stage default não encontrado ou sem handler\n")
		}
		return true
		
	case "link", "acessar", "formulário", "formulario":
		// Mostra o link de acesso
		message := `🔗 *Link para Adesão*

Para acessar o formulário de adesão, clique no link abaixo:

📋 *Formulário de Pessoa Física:*
https://wscredcoopsbf.facilinformatica.com.br/facweb/#formulario-de-pessoa-fisica

💡 *Dica:* Você pode copiar e colar o link no seu navegador.

Digite *0* para voltar ao menu principal.`
		
		m.Reply(message)
		return true
		
	default:
		fmt.Printf("🔄 [ADESAO] Enviando mensagem padrão de adesão\n")
		// Mostra as instruções de adesão
		message := `📋 *PROCESSO DE ADESÃO - ATIVA GRUPO SBF*

Para aderir à Ativa, siga os passos abaixo:

🔗 *1. Acesse o Link*
Para aderir à Ativa, acesse o link:
https://wscredcoopsbf.facilinformatica.com.br/facweb/#formulario-de-pessoa-fisica

📝 *2. Preencha o Formulário*
Em seguida, preencha os campos obrigatórios marcados com asterisco vermelho (*).

💾 *3. Salve os Dados*
Após inserir todos os dados necessários, clique em "SALVAR"

📄 *4. Termo de Consentimento*
Aparecerá na tela o Termo de Consentimento de Alteração de Dados Cadastrais. Leia atentamente e dê o aceite para prosseguir.

✅ *5. Confirmação*
Agora é só aguardar que o nosso time irá enviar um e-mail de boas-vindas e confirmação do cadastro.

💡 *Comandos disponíveis:*
• Digite *link* para acessar o formulário
• Digite *0* para voltar ao menu principal

Precisa de mais alguma informação sobre o processo de adesão?`
		
		m.Reply(message)
		return true
	}
	
	fmt.Printf("⚠️ [ADESAO] Nenhum caso foi executado para: '%s'\n", text)
	return false
}

// Handler do stage de aplicativo/senha
func aplicativoHandler(conn *IClient, m *IMessage, userStage *UserStage) bool {
	fmt.Printf("🚀 [APLICATIVO] Handler INICIADO para usuário %s\n", m.Sender.ToNonAD().User)
	fmt.Printf("🚀 [APLICATIVO] Parâmetros: conn=%v, m=%v, userStage=%v\n", conn != nil, m != nil, userStage != nil)
	
	text := strings.ToLower(strings.TrimSpace(m.Text))
	
	fmt.Printf("🔍 [APLICATIVO] Handler recebeu: '%s' do usuário %s\n", text, m.Sender.ToNonAD().User)
	fmt.Printf("🔍 [APLICATIVO] Texto processado: '%s'\n", text)
	
	switch text {
	case "0", "voltar", "menu", "início", "inicio":
		fmt.Printf("🔄 [APLICATIVO] Usuário quer voltar ao menu principal\n")
		err := ChangeUserStage(m.Sender.ToNonAD().User, "default")
		if err != nil {
			fmt.Printf("❌ [APLICATIVO] Erro ao mudar stage: %s\n", err.Error())
			m.Reply("❌ Erro ao voltar: " + err.Error())
			return false
		}
		fmt.Printf("✅ [APLICATIVO] Stage mudado para 'default'\n")
		defaultStage := GetStage("default")
		if defaultStage != nil && defaultStage.Handler != nil {
			fmt.Printf("🔄 [APLICATIVO] Executando handler do stage default\n")
			userStage, _ := GetUserStage(m.Sender.ToNonAD().User)
			defaultStage.Handler(conn, m, userStage)
			fmt.Printf("✅ [APLICATIVO] Handler do default executado\n")
		} else {
			fmt.Printf("❌ [APLICATIVO] Stage default não encontrado ou sem handler\n")
		}
		return true
		
	case "1", "baixar", "download", "aplicativo":
		fmt.Printf("🔄 [APLICATIVO] Usuário quer saber como baixar o aplicativo\n")
		message := `📱 *Como baixar o aplicativo*

Você pode encontrar o nosso aplicativo pesquisando por "Cooper Ativa" em iOS ou Android.

🔍 *Como encontrar:*
• **iOS (App Store):** Procure por "Cooper Ativa"
• **Android (Google Play):** Procure por "Cooper Ativa"

💡 *Dica:* Certifique-se de baixar o aplicativo oficial da Cooperativa Ativa.

📋 *Navegação:*
• Digite *0* para voltar ao menu principal
• Digite *5* para encerrar atendimento`
		
		m.Reply(message)
		return true
		
	case "2", "esqueci", "senha", "recuperar":
		fmt.Printf("🔄 [APLICATIVO] Usuário quer recuperar senha\n")
		message := `🔑 *Esqueci minha senha de acesso ao aplicativo*

*Siga os passos abaixo:*

1️⃣ **Acesse o iBanking através deste link:**
https://wscredcoopsbf.facilinformatica.com.br/facweb/

2️⃣ **Informe seu CPF e clique no botão "próxima".**

3️⃣ **Clique no botão "esqueceu a senha?"**

4️⃣ **Digite o CPF e a data de nascimento e clique botão "enviar"**

5️⃣ **Você receberá uma senha temporária no e-mail cadastrado na Ativa**

6️⃣ **Após o recebimento, entre no site ou app da Cooper Ativa novamente, repita o passo 1 e entre utilizando a sua senha temporária**

7️⃣ **Após entrar, será necessário criar a sua senha definitiva. Para isso, insira sua senha temporária em "Senha atual", e crie a sua nova senha de 6 dígitos nos demais campos**

8️⃣ **Uma vez confirmada a nova senha definitiva, clique em "ALTERAR SENHA"**

9️⃣ **Para finalizar, aceite o termo de Consentimento para Tratamento de Dados para continuar.**

📋 *Navegação:*
• Digite *0* para voltar ao menu principal
• Digite *5* para encerrar atendimento`
		
		m.Reply(message)
		return true
		
	case "3", "bloqueada", "bloqueado":
		fmt.Printf("🔄 [APLICATIVO] Usuário tem senha bloqueada\n")
		message := `🔒 *Senha bloqueada*

Você tentou realizar o acesso via iBanking ou pelo aplicativo "Cooper Ativa" e recebeu a mensagem que sua senha estava bloqueada? 🔒

**Opções:**
• Digite *1* se SIM
• Digite *2* se NÃO

📋 *Navegação:*
• Digite *0* para voltar ao menu principal
• Digite *5* para encerrar atendimento`
		
		m.Reply(message)
		return true
		
	case "4", "voltar menu", "menu inicial":
		fmt.Printf("🔄 [APLICATIVO] Usuário quer voltar ao menu inicial\n")
		err := ChangeUserStage(m.Sender.ToNonAD().User, "default")
		if err != nil {
			fmt.Printf("❌ [APLICATIVO] Erro ao mudar stage: %s\n", err.Error())
			m.Reply("❌ Erro ao voltar: " + err.Error())
			return false
		}
		fmt.Printf("✅ [APLICATIVO] Stage mudado para 'default'\n")
		defaultStage := GetStage("default")
		if defaultStage != nil && defaultStage.Handler != nil {
			fmt.Printf("🔄 [APLICATIVO] Executando handler do stage default\n")
			userStage, _ := GetUserStage(m.Sender.ToNonAD().User)
			defaultStage.Handler(conn, m, userStage)
			fmt.Printf("✅ [APLICATIVO] Handler do default executado\n")
		} else {
			fmt.Printf("❌ [APLICATIVO] Stage default não encontrado ou sem handler\n")
		}
		return true
		
	case "5", "encerrar", "sair", "fim":
		fmt.Printf("🔄 [APLICATIVO] Usuário quer encerrar atendimento\n")
		message := `👋 *Atendimento encerrado!*

Obrigado por entrar em contato conosco.

Se precisar de mais alguma coisa, é só me chamar novamente! 😊`
		
		m.Reply(message)
		return true
		
	// Sub-opções para senha bloqueada
	case "sim", "1 sim":
		fmt.Printf("🔄 [APLICATIVO] Usuário confirmou que tem senha bloqueada\n")
		message := `📞 *Atendimento para senha bloqueada*

Informe sua matrícula e aguarde um instante, você será atendido em breve.

Nossa equipe entrará em contato com você para solucionar o bloqueio o mais breve possível.

📋 *Navegação:*
• Digite *0* para voltar ao menu principal
• Digite *5* para encerrar atendimento`
		
		m.Reply(message)
		return true
		
	case "não", "nao", "2 não", "2 nao":
		fmt.Printf("🔄 [APLICATIVO] Usuário negou que tem senha bloqueada\n")
		message := `📧 *Reporte o erro*

Envie um print da tela com o erro para o e-mail cooperativa@gruposbf.com.br para que possamos verificar o erro.

Nossa equipe entrará em contato com você para solucionar o bloqueio o mais breve possível.

📋 *Navegação:*
• Digite *0* para voltar ao menu principal
• Digite *5* para encerrar atendimento`
		
		m.Reply(message)
		return true
		
	default:
		fmt.Printf("🔄 [APLICATIVO] Enviando mensagem padrão do aplicativo\n")
		// Mostra o menu do aplicativo/senha
		message := `📱 *APLICATIVO OU SENHA DE ACESSO*

Escolha a opção desejada:

1️⃣ *Como baixar o aplicativo*
2️⃣ *Esqueci minha senha de acesso ao aplicativo*
3️⃣ *Senha bloqueada*
4️⃣ *Voltar ao menu inicial*
5️⃣ *Encerrar atendimento*

💡 *Como usar:*
• Digite o *número* da opção (ex: 1, 2, 3...)
• Digite palavras-chave como *baixar*, *senha*, *bloqueada*

Escolha uma opção para continuar! ⬇️`
		
		m.Reply(message)
		return true
	}
	
	fmt.Printf("⚠️ [APLICATIVO] Nenhum caso foi executado para: '%s'\n", text)
	return false
}

// Registra um novo stage
func RegisterStage(stage *Stage) {
	if stages == nil {
		stages = make(map[string]*Stage)
	}
	stages[stage.ID] = stage
}

// Obtém um stage por ID
func GetStage(id string) *Stage {
	return stages[id]
}

// Obtém todos os stages registrados
func GetAllStages() map[string]*Stage {
	return stages
}

// Obtém o stage atual do usuário
func GetUserStage(userID string) (*UserStage, error) {
	query := "SELECT user_id, current_stage, data, created_at, updated_at FROM user_stages WHERE user_id = ?"
	row := db.QueryRow(query, userID)
	
	var userStage UserStage
	var dataJSON string
	
	err := row.Scan(&userStage.UserID, &userStage.CurrentStage, &dataJSON, &userStage.CreatedAt, &userStage.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
		// Usuário não existe, retorna stage padrão
		return &UserStage{
			UserID:      userID,
			CurrentStage: "default",
			Data:        make(map[string]interface{}),
			CreatedAt:   time.Now().Unix(),
			UpdatedAt:   time.Now().Unix(),
		}, nil
		}
		return nil, err
	}
	
	// Deserializa os dados JSON
	if dataJSON != "" {
		err = json.Unmarshal([]byte(dataJSON), &userStage.Data)
		if err != nil {
			userStage.Data = make(map[string]interface{})
		}
	} else {
		userStage.Data = make(map[string]interface{})
	}
	
	return &userStage, nil
}

// Salva ou atualiza o stage do usuário
func SaveUserStage(userStage *UserStage) error {
	dataJSON, err := json.Marshal(userStage.Data)
	if err != nil {
		return err
	}
	
	now := time.Now().Unix()
	userStage.UpdatedAt = now
	
	query := `
	INSERT OR REPLACE INTO user_stages (user_id, current_stage, data, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)`
	
	_, err = db.Exec(query, userStage.UserID, userStage.CurrentStage, string(dataJSON), userStage.CreatedAt, userStage.UpdatedAt)
	return err
}

// Muda o usuário para um novo stage
func ChangeUserStage(userID string, newStageID string) error {
	return ChangeUserStageWithMessage(userID, newStageID, nil, nil)
}

// ChangeUserStageWithMessage muda o stage do usuário e opcionalmente executa o handler
func ChangeUserStageWithMessage(userID string, newStageID string, conn *IClient, m *IMessage) error {
	userStage, err := GetUserStage(userID)
	if err != nil {
		return err
	}
	
	// Verifica se o stage existe
	stage := GetStage(newStageID)
	if stage == nil {
		return fmt.Errorf("stage '%s' não encontrado", newStageID)
	}
	
	// Verifica se o usuário pode acessar este stage
	if stage.IsOwner && !isOwner(userID) {
		return fmt.Errorf("você não tem permissão para acessar este stage")
	}
	
	userStage.CurrentStage = newStageID
	userStage.Data = make(map[string]interface{}) // Limpa dados do stage anterior
	
	err = SaveUserStage(userStage)
	if err != nil {
		return err
	}
	
	// Se foi fornecido conn e m, executa o handler do novo stage
	if conn != nil && m != nil && stage.Handler != nil {
		// Executa o handler do novo stage diretamente
		stage.Handler(conn, m, userStage)
	}
	
	return nil
}

// Verifica se o usuário pode navegar para um stage específico
func CanNavigateToStage(userID string, fromStageID string, toStageID string) bool {
	fromStage := GetStage(fromStageID)
	if fromStage == nil {
		return false
	}
	
	// Verifica se o stage de destino está na lista de próximos stages
	for _, nextStage := range fromStage.NextStages {
		if nextStage == toStageID {
			return true
		}
	}
	
	// Permite navegação para o próprio stage
	return fromStageID == toStageID
}

// Processa uma mensagem no contexto do stage atual do usuário
func ProcessStageMessage(conn *IClient, m *IMessage) bool {
	userID := m.Sender.ToNonAD().User
	
	fmt.Printf("🔍 [STAGES] Processando mensagem '%s' do usuário %s\n", m.Text, userID)
	
	// Verifica se o usuário está autorizado (apenas 5514991983652)
	authorizedNumber := "5514991983652"
	if userID != authorizedNumber {
		fmt.Printf("❌ [STAGES] Usuário não autorizado: %s\n", userID)
		m.Reply("❌ *Acesso não autorizado*\n\nEste atendimento é restrito a usuários específicos.\n\nSe você acredita que deveria ter acesso, entre em contato com a administração.")
		return false
	}
	
	// Obtém o stage atual do usuário
	userStage, err := GetUserStage(userID)
	if err != nil {
		fmt.Printf("❌ [STAGES] Erro ao obter stage do usuário: %s\n", err.Error())
		m.Reply("Erro ao obter informações do usuário: " + err.Error())
		return false
	}
	
	fmt.Printf("🔍 [STAGES] Usuário está no stage: %s\n", userStage.CurrentStage)
	
	// Obtém o stage atual
	stage := GetStage(userStage.CurrentStage)
	if stage == nil {
		// Se o stage não existe, volta para o default
		userStage.CurrentStage = "default"
		SaveUserStage(userStage)
		stage = GetStage("default")
		if stage == nil {
			m.Reply("Sistema de stages não inicializado corretamente.")
			return false
		}
	}
	 
	// Verifica permissões do stage
	if stage.IsOwner && !m.IsOwner {
		m.Reply("Você não tem permissão para acessar este stage.")
		return false
	}
	
	if stage.IsGroup && !m.Info.IsGroup {
		m.Reply("Este stage só funciona em grupos.")
		return false
	}
	
	if stage.IsPrivate && m.Info.IsGroup {
		m.Reply("Este stage só funciona em conversas privadas.")
		return false
	}
	
	// Executa o handler do stage
	if stage.Handler != nil {
		fmt.Printf("🔄 [STAGES] Executando handler do stage '%s'\n", stage.ID)
		fmt.Printf("🔄 [STAGES] Handler function: %v\n", stage.Handler)
		fmt.Printf("🔄 [STAGES] Chamando handler...\n")
		result := stage.Handler(conn, m, userStage)
		fmt.Printf("✅ [STAGES] Handler executado, resultado: %v\n", result)
		return result
	} else {
		fmt.Printf("❌ [STAGES] Stage sem handler\n")
	}
	
	return false
}

// Função auxiliar para verificar se é owner
func isOwner(userID string) bool {
	owners := os.Getenv("OWNER")
	if owners == "" {
		return false
	}
	
	// Verifica se o userID está na lista de owners (separados por vírgula)
	ownerList := strings.Split(owners, ",")
	for _, owner := range ownerList {
		if strings.TrimSpace(owner) == userID {
			return true
		}
	}
	return false
}

// Fecha a conexão com o banco de dados
func CloseStagesDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

