package botcreateauth

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DebugHandling(err error, update tgbotapi.Update, bot *tgbotapi.BotAPI) (bool){
	var msgErr string 
	var debugMsg string
	if err != nil {
		log.Printf(debugMsg, err)
		ErrorHandling(update, bot, msgErr)
		return true
	}
	return false
}

// "[DEBUG] Error detecting language"