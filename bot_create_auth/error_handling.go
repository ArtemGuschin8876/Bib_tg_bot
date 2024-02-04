package botcreateauth

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func ErrorHandling(update tgbotapi.Update, bot *tgbotapi.BotAPI, messageErr string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageErr)
	bot.Send(msg)
}






