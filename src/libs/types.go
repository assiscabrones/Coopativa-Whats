package libs

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

type IClient struct {
	WA *whatsmeow.Client
}

type ICommand struct {
	Name        string
	As          []string
	Description string
	Tags        string
	IsPrefix    bool
	IsOwner     bool
	IsMedia     bool
	IsQuery     bool
	IsGroup     bool
	IsWait      bool
	IsPrivate   bool
	Before      func(conn *IClient, m *IMessage)
	Execute     func(conn *IClient, m *IMessage) bool
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
