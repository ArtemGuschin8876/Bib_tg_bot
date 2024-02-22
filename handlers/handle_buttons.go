package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Создание главного меню, здесь все кнопки которые будут добавлены в Бота.
func createMainMenu() tgbotapi.InlineKeyboardMarkup {
	trButton := tgbotapi.NewInlineKeyboardButtonData("Перевод", "translate")
	helpButton := tgbotapi.NewInlineKeyboardButtonData("Инфо", "information")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(helpButton),
		tgbotapi.NewInlineKeyboardRow(trButton),
	)

	return keyboard
}

// Обработка созданных в createMainMenu кнопок и их реализация
func HandleButton(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	text := update.Message.Text
	var msgText string

	if text == "start" {
		msgText = "Выберите что вас интересует!"
	} else if text != "" {
		msgText = "Я вас не понимаю, пожалуйста воспользуйтесь меню:)"
	}

	sendMessageWithButtons(bot, update.Message.Chat.ID, msgText, createMainMenu())

}

// Отправляет сообщение с кнопками
func sendMessageWithButtons(bot *tgbotapi.BotAPI, chatID int64, text string, replyMarkup tgbotapi.InlineKeyboardMarkup) error {
	if text == "" {
		log.Println("Error: Attempt to send a message with empty text")
	}
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = replyMarkup

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
	return nil
}

// Функция добавляет кнопки "Продолжить" или "Закончить" перевод текста.
func SendMessageWithContinueAndFinishButton(bot *tgbotapi.BotAPI, chatID int64, text string) {

	if text == "" {
		log.Println("Error: Attempt to send a message with empty text")
	}
	msg := tgbotapi.NewMessage(chatID, text)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Продолжить", "continue_translation"),
			tgbotapi.NewInlineKeyboardButtonData("Закончить", "finish_translation"),
		),
	)

	msg.ReplyMarkup = inlineKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
