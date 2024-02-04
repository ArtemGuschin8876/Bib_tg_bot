package botcreateauth

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	msgErr   string
	debugMsg string
)

func DebugHandling(err error, update tgbotapi.Update, bot *tgbotapi.BotAPI) bool {

	if err != nil {
		log.Printf(debugMsg, err)
		ErrorHandling(update, bot, msgErr)
		return true
	}

	return false
}

func ErrorHandling(update tgbotapi.Update, bot *tgbotapi.BotAPI, messageErr string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageErr)
	bot.Send(msg)
}