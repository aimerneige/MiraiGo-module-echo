package echo

import (
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

var instance *echo
var logger = utils.GetModuleLogger("internal.logging")

type echo struct {
}

func init() {
	instance = &echo{}
	bot.RegisterModule(instance)
}

func (e *echo) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "aimerneige.test.echo",
		Instance: instance,
	}
}

func (m *echo) Init() {
}

func (m *echo) PostInit() {
}

func (m *echo) Serve(b *bot.Bot) {
	b.OnGroupMessage(func(c *client.QQClient, msg *message.GroupMessage) {
		out := echoMessage(msg.ToString())
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendGroupMessage(msg.GroupCode, m)
	})

	b.OnPrivateMessage(func(c *client.QQClient, msg *message.PrivateMessage) {
		out := echoMessage(msg.ToString())
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendPrivateMessage(msg.Sender.Uin, m)
	})
}

func (m *echo) Start(b *bot.Bot) {
}

func (m *echo) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

func echoMessage(msg string) string {
	if msg[:5] == "echo " {
		return msg[5:]
	}
	return ""
}
