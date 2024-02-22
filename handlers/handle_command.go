package handlers

import (
	"errors"
	"fmt"
	"log"
	"projects/BIb_bot/translate"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	errMsgEmptyText      = "Пожалуйста, введите текст для перевода с командой /tr."
	errMsgDetectLanguage = "Произошла ошибка при определении языка текста."
	errMsgTranslation    = "Произошла ошибка при переводе текста."
)

// Обработчик команд, возможно эта функция уберётся
func HandleCommand(update tgbotapi.Update) string {
	command := update.Message.Command()

	switch command {
	case "todo":
		return "Здесь весь лист"
	case "help":
		return "Все команды:\n /tr - Перевод языка ru <-> en.\n "
	case "start":
		return "Привет я BiB, воспользуйся командой /help и узнай о моих возможностях."
	default:
		return "Default"
	}
}

// Определяет Введенный текст, и переводит его  en<->ru
func HandleTranslateCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI, userStates map[int64]UserState) (string, error) {
	text := strings.TrimSpace(update.Message.Text)
	chatID := update.Message.Chat.ID

	if text == "" {
		return "", errors.New(errMsgEmptyText)
	}

	detectLanguage, err := translate.DetectLanguage(text)
	if err != nil {
		log.Println("[DEBUG] Ошибка при определении языка: ", err)
		return "", errors.New(errMsgDetectLanguage)
	}

	sourceLanguage := detectLanguage.Language.String()

	targetLanguage := translate.DetermineTargetLanguage(sourceLanguage)

	translatedText, err := translate.TranslateTextWithModel(targetLanguage, text, "nmt")
	if err != nil {

		log.Println("[DEBUG] Ошибка при переводе текста: ", err)
		errorrMessage := "Произошла ошибка при определении языка текста."
		SendMessageWithContinueAndFinishButton(bot, chatID, errorrMessage)

		return "", err
	}

	log.Println("[DEBUG] Переведённый текст: 	", translatedText)

	translatedText = fmt.Sprintf("Перевод: \n%s", translatedText)
	SendMessageWithContinueAndFinishButton(bot, chatID, translatedText)
	// userStates[chatID] = UserState{TranslationPrompt: true}

	return translatedText, nil

}
