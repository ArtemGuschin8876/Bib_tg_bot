package handlers

import (
	botcommands "projects/BIb_bot/bot_commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(update tgbotapi.Update) string {
	command := update.Message.Command()

	switch command {
	case "todo" :
		return "Здесь весь лист"
	case "help" :
		return "Тут помощиии"
	case "tr" :
		return botcommands.HandleTranslateCommand(update)
	default :
		return "Ничего не происходит"
	}
}

