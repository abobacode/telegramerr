package telegramerr

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewBot(token string, chatID int64) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	botAPI.Debug = false

	return &Bot{
		api:    botAPI,
		chatID: chatID,
	}, nil
}
