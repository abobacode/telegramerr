package telegramerr

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Bot struct {
	api    *tgbotapi.BotAPI
	chatID int64
}
