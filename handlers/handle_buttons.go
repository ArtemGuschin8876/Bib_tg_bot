package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	msgText string
)

func createMainMenu() tgbotapi.InlineKeyboardMarkup {
	trButton := tgbotapi.NewInlineKeyboardButtonData("Переводчик", "/tr")
	helpButton := tgbotapi.NewInlineKeyboardButtonData("Помощь", "/help")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(helpButton),
		tgbotapi.NewInlineKeyboardRow(trButton),
	)

	return keyboard
}

func HandleButton(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	text := update.Message.Text

	if text == "start" {

		msgText = "Выберите что вас интересует!"
		sendMessageWithButtons(bot, update.Message.Chat.ID, msgText, createMainMenu())

	} else if text != "" {

		msgText = "Я вас не понимаю, пожалуйста воспользуйтесь меню:)"
		sendMessageWithButtons(bot, update.Message.Chat.ID, msgText, createMainMenu())

	}

}

func sendMessageWithButtons(bot *tgbotapi.BotAPI, chatID int64, text string, replyMarkup tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = replyMarkup

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
