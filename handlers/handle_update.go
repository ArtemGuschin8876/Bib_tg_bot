package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	msgText string
)

func HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message == nil {
		return
	}
	if update.Message.IsCommand() {
		msgText = HandleCommand(update)
	} else {
		msgText = "Ничего не происходит"
	}

	SendMessage(update.Message.Chat.ID, msgText, bot)
}
