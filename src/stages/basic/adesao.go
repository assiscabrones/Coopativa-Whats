package basic

import (
	"fmt"
	"strings"
	"hisoka/src/libs"
)

func init() {
	// Registra o stage de adesão
	libs.RegisterStage(&libs.Stage{
		ID:          "adesao",
		Name:        "Adesão",
		Description: "Processo de adesão à Ativa Grupo SBF",
		Handler:     adesaoHandler,
		NextStages:  []string{"default"},
		IsOwner:     false,
		IsGroup:     false,
		IsPrivate:   false,
	})
}

// Handler do stage de adesão
func adesaoHandler(conn *libs.IClient, m *libs.IMessage, userStage *libs.UserStage) bool {
	fmt.Printf("🚀 [ADESAO] Handler INICIADO para usuário %s\n", m.Sender.ToNonAD().User)
	fmt.Printf("🚀 [ADESAO] Parâmetros: conn=%v, m=%v, userStage=%v\n", conn != nil, m != nil, userStage != nil)
	
	text := strings.ToLower(strings.TrimSpace(m.Text))
	
	fmt.Printf("🔍 [ADESAO] Handler recebeu: '%s' do usuário %s\n", text, m.Sender.ToNonAD().User)
	fmt.Printf("🔍 [ADESAO] Texto processado: '%s'\n", text)
	
	switch text {
	case "0", "voltar", "menu", "início", "inicio":
		fmt.Printf("🔄 [ADESAO] Usuário quer voltar ao menu principal\n")
		
		// Volta para o menu principal
		err := libs.ChangeUserStage(m.Sender.ToNonAD().User, "default")
		if err != nil {
			fmt.Printf("❌ [ADESAO] Erro ao mudar stage: %s\n", err.Error())
			m.Reply("❌ Erro ao voltar: " + err.Error())
			return false
		}
		
		fmt.Printf("✅ [ADESAO] Stage mudado para 'default'\n")
		
		// Executa o handler do stage default diretamente
		defaultStage := libs.GetStage("default")
		if defaultStage != nil && defaultStage.Handler != nil {
			fmt.Printf("🔄 [ADESAO] Executando handler do stage default\n")
			userStage, _ := libs.GetUserStage(m.Sender.ToNonAD().User)
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
	
	// Se chegou aqui, nenhum caso foi executado
	fmt.Printf("⚠️ [ADESAO] Nenhum caso foi executado para: '%s'\n", text)
	return false
}
