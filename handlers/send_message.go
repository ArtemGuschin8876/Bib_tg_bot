package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(userID int64, msgText string, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(userID, msgText)

	_, err := bot.Send(msg)

	if err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
	}

}
