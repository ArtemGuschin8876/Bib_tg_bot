package handlers

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserState struct {
	TranslationPrompt bool
	AdminDataPrompt   bool
}

const (
	callbackInfo      = "information"
	callbackTranslate = "translate"
	callbackContinue  = "continue_translation"
	callbackFinish    = "finish_translation"
	callbackAdmin     = "admin_data"
)

// Главный обработчик (команды, кнопки, состояния)
func HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI, userStates map[int64]UserState) {

	if update.Message == nil {

		if update.CallbackQuery != nil {

			handleCallback(bot, update.CallbackQuery, userStates, update)
		}

		return

	}

	if HandleTranslationPrompt(update, bot, userStates) {
		return;
	}

	if update.Message.IsCommand() {

		msgText := HandleCommand(update)
		SendMessageWithoutButtons(bot, update.Message.Chat.ID, msgText)

	} else {

		HandleButton(update, bot)

	}

}

// Отправляет сообщение без добавления кнопок
func SendMessageWithoutButtons(bot *tgbotapi.BotAPI, chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
	return nil
}

// Обработка ответа от всех кнопок
func handleCallback(bot *tgbotapi.BotAPI, query *tgbotapi.CallbackQuery, userStates map[int64]UserState, update tgbotapi.Update) {
	log.Printf("Received callback: %+v", query)

	switch query.Data {

	case callbackInfo:
		msgText := "Здесь все команды"

		if err := SendMessageWithoutButtons(bot, query.Message.Chat.ID, msgText); err != nil {
			log.Printf("Error sending message: %v", err)
		}

	case callbackTranslate, callbackContinue:
		msgText := "Введите текст для перевода:"

		userStates[query.Message.Chat.ID] = UserState{TranslationPrompt: true}

		if err := SendMessageWithoutButtons(bot, query.Message.Chat.ID, msgText); err != nil {
			log.Printf("Error sending message: %v", err)
		}

	case callbackFinish:
		msgText := "Меню"
		userID := query.Message.From.ID

		if err := sendMessageWithButtons(bot, query.Message.Chat.ID, msgText, createMainMenu(int(userID))); err != nil {
			log.Printf("Error sending message: %v", err)
		}

	case callbackAdmin:
		msgText := "Данные только для администратора"
		userID := query.Message.From.ID

		userStates[query.Message.Chat.ID] = UserState{TranslationPrompt: true}

		if err := sendMessageWithButtons(bot, query.Message.Chat.ID, msgText, createMainMenu(int(userID))); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

func HandleTranslationPrompt(update tgbotapi.Update, bot *tgbotapi.BotAPI, userStates map[int64]UserState) bool {
	if state, ok := userStates[update.Message.Chat.ID]; ok {
		if state.TranslationPrompt {

			chatID := update.Message.Chat.ID

			_, err := HandleTranslateCommand(update, bot, userStates)
			if err != nil {
				fmt.Println("Error")
			}

			userStates[chatID] = UserState{TranslationPrompt: false}
			return true;
		}
		return false;
	} else {
		return false;
	}
}
