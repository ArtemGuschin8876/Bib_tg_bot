package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserState struct {
	TranslationPrompt bool
}

func HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI, userStates map[int64]UserState) {

	if update.Message == nil {

		if update.CallbackQuery != nil {

			handleCallback(bot, update.CallbackQuery, update, userStates)
		}

		return

	}

	if state, ok := userStates[update.Message.Chat.ID]; ok {
		if state.TranslationPrompt {

			chatID := update.Message.Chat.ID

			translatedText := HandleTranslateCommand(update, bot, userStates)

			SendMessageWithoutButtons(bot, chatID, translatedText)

			userStates[chatID] = UserState{TranslationPrompt: false}
			return
		}
	}

	if update.Message.IsCommand() {

		msgText = HandleCommand(update)
		SendMessageWithoutButtons(bot, update.Message.Chat.ID, msgText)

	} else {

		HandleButton(update, bot)

	}

}

func SendMessageWithoutButtons(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func handleCallback(bot *tgbotapi.BotAPI, query *tgbotapi.CallbackQuery, update tgbotapi.Update, userStates map[int64]UserState) {
	log.Printf("Received callback: %+v", query)

	// chatID := update.Message.Chat.ID

	switch query.Data {
	case "help":
		msgText := "Здесь все команды"
		SendMessageWithoutButtons(bot, query.Message.Chat.ID, msgText)
	case "tr":
		userStates[query.Message.Chat.ID] = UserState{TranslationPrompt: true}
		msgText := "Введите текст для перевода:"
		SendMessageWithoutButtons(bot, query.Message.Chat.ID, msgText)

	case "continue_translation":
		userStates[query.Message.Chat.ID] = UserState{TranslationPrompt: true}
		msgText := "Введите текст для перевода:"
		SendMessageWithoutButtons(bot, query.Message.Chat.ID, msgText)
	case "finish_translation":
		msgText = "Меню"
		sendMessageWithButtons(bot, query.Message.Chat.ID, msgText, createMainMenu())
	}
}
