package telegramerr

import (
	"fmt"
	"log"
	time2 "time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const ERROR = "🔴 ERROR: %s" +
	"\n" + "⚙️ SERVICE: %s" +
	"\n" + "🕘 TIME: %s" +
	"\n" + "⛓️ TRACING: //TODO"

func (b *Bot) SendErrorMessage(err error, service string) {
	time := time2.Now()
	msg := tgbotapi.NewMessage(b.chatID, fmt.Sprintf(ERROR, err, service, time))
	if _, err := b.api.Send(msg); err != nil {
		log.Printf("Failed to send error message: %v", err)
	}
}
