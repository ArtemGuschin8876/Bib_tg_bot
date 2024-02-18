package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {

		msgText = HandleCommand(update)
		sendMessageWithoutButtons(bot, update.Message.Chat.ID, msgText)
	} else {

		HandleButton(update, bot)

	}

}

func sendMessageWithoutButtons(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
