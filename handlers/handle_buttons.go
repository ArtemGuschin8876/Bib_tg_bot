package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	msgText string
)

func CreateMainMenu() tgbotapi.ReplyKeyboardMarkup {

	helpButton := tgbotapi.NewKeyboardButton("/help")

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(helpButton),
	)

	return keyboard
}

func HandleButton(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	switch update.Message.Text {
	case "start":
		msgText = "Menu"
		sendMessageWithButtons(bot, update.Message.Chat.ID, msgText, CreateMainMenu())
	default:
		msgText = "Тут дефолт"
		sendMessageWithButtons(bot, update.Message.Chat.ID, msgText, CreateMainMenu())

	}

}

func sendMessageWithButtons(bot *tgbotapi.BotAPI, chatID int64, text string, replyMarkup tgbotapi.ReplyKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = replyMarkup

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
