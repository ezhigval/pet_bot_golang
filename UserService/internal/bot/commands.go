package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func (b *Bot) startCommand(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(
		msg.Chat.ID,
		"Привет! Я учебный Go-бот 🤖\nПока я умею только эхо, но это ненадолго.",
	)
	b.api.Send(reply)
}

func (b *Bot) infoCommand(msg *tgbotapi.Message) {
	start := time.Now()

	user := msg.From
	chat := msg.Chat

	serverTime := time.Now().UTC()
	latency := time.Since(start)

	text := fmt.Sprintf(
		"ℹ️ *Bot Info*\n\n"+
			"👤 User ID: `%d`\n"+
			"🧑 Username: `%s`\n"+
			"💬 Chat ID: `%d`\n"+
			"📨 Chat type: `%s`\n\n"+
			"🕒 Server time (UTC): `%s`\n"+
			"⚡ Response time: `%s`\n"+
			"🤖 Bot: `@%s`",
		user.ID,
		nullable(user.UserName),
		chat.ID,
		chat.Type,
		serverTime.Format(time.RFC3339),
		latency,
		b.api.Self.UserName,
	)

	reply := tgbotapi.NewMessage(chat.ID, text)
	reply.ParseMode = "Markdown"

	b.api.Send(reply)
}

func (b *Bot) unknownCommand(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(
		msg.Chat.ID,
		"Неизвестная команда 🤔",
	)
	b.api.Send(reply)
}

func (b *Bot) echo(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(
		msg.Chat.ID,
		msg.Text,
	)
	b.api.Send(reply)
}

func nullable(value string) string {
	if value == "" {
		return "—"
	}
	return value
}
