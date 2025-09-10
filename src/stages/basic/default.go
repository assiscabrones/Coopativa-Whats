package basic

import (
	"fmt"
	"strings"
	"hisoka/src/libs"
)

func init() {
	// Registra o stage padrão
	libs.RegisterStage(&libs.Stage{
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
}

// Handler do stage default
func defaultHandler(conn *libs.IClient, m *libs.IMessage, userStage *libs.UserStage) bool {
	// Teste simples primeiro
	fmt.Println("🚀 [DEFAULT] TESTE - Handler INICIADO!")
	
	fmt.Printf("🚀 [DEFAULT] Handler INICIADO para usuário %s\n", m.Sender.ToNonAD().User)
	fmt.Printf("🚀 [DEFAULT] Parâmetros: conn=%v, m=%v, userStage=%v\n", conn != nil, m != nil, userStage != nil)
	
	text := strings.ToLower(strings.TrimSpace(m.Text))

	fmt.Printf("🔍 [DEFAULT] Handler recebeu: '%s' do usuário %s\n", text, m.Sender.ToNonAD().User)

	// Processa as opções do menu
	switch text {
	case "1", "adesão", "adesao":
		fmt.Printf("🔄 [DEFAULT] Usuário quer ir para adesão\n")
		
		// Navega para stage de adesão
		err := libs.ChangeUserStage(m.Sender.ToNonAD().User, "adesao")
		if err != nil {
			fmt.Printf("❌ [DEFAULT] Erro ao mudar stage: %s\n", err.Error())
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		
		fmt.Printf("✅ [DEFAULT] Stage mudado para 'adesao'\n")
		
		// Executa o handler do stage de adesão diretamente
		adesaoStage := libs.GetStage("adesao")
		if adesaoStage != nil && adesaoStage.Handler != nil {
			fmt.Printf("🔄 [DEFAULT] Executando handler do stage adesao\n")
			userStage, _ := libs.GetUserStage(m.Sender.ToNonAD().User)
			adesaoStage.Handler(conn, m, userStage)
			fmt.Printf("✅ [DEFAULT] Handler do adesao executado\n")
		} else {
			fmt.Printf("❌ [DEFAULT] Stage adesao não encontrado ou sem handler\n")
		}
		return true

	case "2", "aplicativo", "senha", "acesso":
		// Navega para stage de aplicativo/senha
		err := libs.ChangeUserStageWithMessage(m.Sender.ToNonAD().User, "aplicativo", conn, m)
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "3", "capital", "investimento":
		// Navega para stage de capital
		err := libs.ChangeUserStageWithMessage(m.Sender.ToNonAD().User, "capital", conn, m)
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "4", "empréstimos", "emprestimos":
		// Navega para stage de empréstimos
		err := libs.ChangeUserStageWithMessage(m.Sender.ToNonAD().User, "emprestimos", conn, m)
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "5", "parcerias":
		// Navega para stage de parcerias
		err := libs.ChangeUserStageWithMessage(m.Sender.ToNonAD().User, "parcerias", conn, m)
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "6", "consultoria", "financeira":
		// Navega para stage de consultoria
		err := libs.ChangeUserStageWithMessage(m.Sender.ToNonAD().User, "consultoria", conn, m)
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "7", "ex-colaborador", "excolaborador":
		// Navega para stage de ex-colaborador
		err := libs.ChangeUserStageWithMessage(m.Sender.ToNonAD().User, "excolaborador", conn, m)
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "8", "negociação", "negociacao", "dívidas", "dividas":
		// Navega para stage de negociação de dívidas
		err := libs.ChangeUserStageWithMessage(m.Sender.ToNonAD().User, "negociacao", conn, m)
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "9", "informe", "rendimentos":
		// Navega para stage de informe de rendimentos
		err := libs.ChangeUserStageWithMessage(m.Sender.ToNonAD().User, "informe", conn, m)
		if err != nil {
			m.Reply("❌ Erro ao acessar: " + err.Error())
			return false
		}
		return true

	case "10", "dúvida", "duvida", "não encontrou", "nao encontrou":
		// Navega para stage de dúvidas
		err := libs.ChangeUserStageWithMessage(m.Sender.ToNonAD().User, "duvidas", conn, m)
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
	
	// Se chegou aqui, nenhum caso foi executado
	fmt.Printf("⚠️ [DEFAULT] Nenhum caso foi executado para: '%s'\n", text)
	return false
}
