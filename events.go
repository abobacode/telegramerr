package telegramerr

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) SendErrorMessage(err error) {
	msg := tgbotapi.NewMessage(b.chatID, fmt.Sprintf("🔴 ERROR: %s", err))
	if _, err := b.api.Send(msg); err != nil {
		log.Printf("Failed to send error message: %v", err)
	}
}
