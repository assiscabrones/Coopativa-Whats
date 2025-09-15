package libs

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

type IClient struct {
	WA *whatsmeow.Client
}

// Estruturas do sistema de stages
type Stage struct {
	ID          string
	Name        string
	Description string
	Handler     func(conn *IClient, m *IMessage, userStage *UserStage) bool
	NextStages  []string // IDs dos stages que podem ser acessados a partir deste
	IsOwner     bool     // Se apenas owners podem acessar
	IsGroup     bool     // Se funciona apenas em grupos
	IsPrivate   bool     // Se funciona apenas em privado
}

type UserStage struct {
	UserID       string
	CurrentStage string
	Data         map[string]interface{} // Dados específicos do usuário no stage atual
	CreatedAt    int64
	UpdatedAt    int64
	LastActivity int64 // Timestamp da última atividade do usuário
}

type IMessage struct {
	Info       types.MessageInfo
	Sender     types.JID
	IsOwner    bool
	Body       string
	Text       string
	Args       []string
	Command    string
	Message    *waE2E.Message
	Media      whatsmeow.DownloadableMessage
	IsMedia    string
	Expiration uint32
	Quoted     *waE2E.ContextInfo
	Reply      func(text string, opts ...whatsmeow.SendRequestExtra) (whatsmeow.SendResponse, error)
	React      func(emoji string, opts ...whatsmeow.SendRequestExtra) (whatsmeow.SendResponse, error)
}
