package handlers

import (
	"context"
	"fmt"
	"hisoka/src/libs"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type IHandler struct {
	Container *store.Device
}

// Timestamp de quando o bot foi inicializado
var botStartupTime = time.Now()

func NewHandler(container *sqlstore.Container) *IHandler {
	// Define o timestamp de inicialização do bot
	botStartupTime = time.Now()
	
	ctx := context.Background()
	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}
	return &IHandler{
		Container: deviceStore,
	}
}

func (h *IHandler) Client() *whatsmeow.Client {
	clientLog := waLog.Stdout("lient", "ERROR", true)
	conn := whatsmeow.NewClient(h.Container, clientLog)
	conn.AddEventHandler(h.RegisterHandler(conn))
	return conn
}

func (h *IHandler) RegisterHandler(conn *whatsmeow.Client) func(evt interface{}) {
	return func(evt interface{}) {
		sock := libs.SerializeClient(conn)
		switch v := evt.(type) {
		case *events.Message:
			m := libs.SerializeMessage(v, sock)

			// skip deleted message
			if m.Message.GetProtocolMessage() != nil && m.Message.GetProtocolMessage().GetType() == 0 {
				return
			}

			// Filtra mensagens antigas (enviadas antes do bot estar online)
			messageTime := v.Info.Timestamp
			if messageTime.Before(botStartupTime) {
				fmt.Printf("\x1b[90m[IGNORADA] Mensagem antiga de %s (%s) - enviada em %s\x1b[39m\n", 
					v.Info.PushName, v.Info.Sender.User, messageTime.Format("15:04:05"))
				return
			}

			// log
			if m.Body != "" {
				fmt.Println("\x1b[94mFrom :", v.Info.PushName, m.Info.Sender.User, "\x1b[39m")
				if len(m.Body) < 350 {
					fmt.Print("\x1b[92mMessage : ", m.Body, "\x1b[39m", "\n")
				} else {
					fmt.Print("\x1b[92mMessage : ", m.Info.Type, "\x1b[39m", "\n")
				}
			}

			// Process stage message
			go ProcessStageMessage(sock, m)
			return
		case *events.Connected, *events.PushNameSetting:
			if len(conn.Store.PushName) == 0 {
				return
			}
			_ = conn.SendPresence(types.PresenceAvailable)
		}
	}
}

func ProcessStageMessage(c *libs.IClient, m *libs.IMessage) {
	// Processa a mensagem usando o sistema de stages
	libs.ProcessStageMessage(c, m)
}

// GetBotStartupTime retorna o timestamp de quando o bot foi inicializado
func GetBotStartupTime() time.Time {
	return botStartupTime
}
