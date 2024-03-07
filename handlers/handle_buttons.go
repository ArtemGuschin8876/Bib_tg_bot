package handlers

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	trButton   = tgbotapi.NewInlineKeyboardButtonData("Перевод", "translate")
	helpButton = tgbotapi.NewInlineKeyboardButtonData("Инфо", "information")
	noteButton = tgbotapi.NewInlineKeyboardButtonData("Админ", "admin_data")
)

// Создание главного меню, здесь все кнопки которые будут добавлены в Бота.
// В зависимости от ID пользователя (админ или нет) выдаются разные клавиатуры.
func createMainMenu(userID int) tgbotapi.InlineKeyboardMarkup {
	adminID, _ := strconv.Atoi(os.Getenv("ADMIN_ID"))

	keyboardWithoutAdmin := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(helpButton),
		tgbotapi.NewInlineKeyboardRow(trButton),
	)

	if userID == adminID {
		keyboardWithoutAdmin.InlineKeyboard = append(keyboardWithoutAdmin.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(noteButton))
	}

	return keyboardWithoutAdmin
}

// Обработка созданных в createMainMenu кнопок и их реализация
func HandleButton(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	userID := update.Message.From.ID

	text := update.Message.Text
	var msgText string

	if text == "start" {
		msgText = "Выберите что вас интересует!"
	} else if text != "" {
		msgText = "Я вас не понимаю, пожалуйста воспользуйтесь меню:)"
	}

	sendMessageWithButtons(bot, update.Message.Chat.ID, msgText, createMainMenu(int(userID)))

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
