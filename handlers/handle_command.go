package handlers

import (
	"fmt"
	"log"
	"projects/BIb_bot/translate"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(update tgbotapi.Update) string {
	command := update.Message.Command()

	switch command {
	case "todo":
		return "Здесь весь лист"
	case "help":
		return "Все команды:\n /tr - Перевод языка ru <-> en.\n "
	// case "tr":
	// return HandleTranslateCommand(update)
	case "start":
		return "Привет я BiB, воспользуйся командой /help и узнай о моих возможностях."
	default:
		return "Default"
	}
}

func HandleTranslateCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI, userStates map[int64]UserState) string {
	// text := strings.TrimSpace(update.Message.CommandArguments())
	text := strings.TrimSpace(update.Message.Text)
	chatID := update.Message.Chat.ID

	if text == "" {
		return "Пожалуйста, введите текст для перевода с командой /tr."
	}

	detectLanguage, err := translate.DetectLanguage(text)

	if err != nil {
		log.Printf("[DEBUG] Ошибка при определении языка: %v", err)

		return "Произошла ошибка при определении языка текста."
	}

	sourceLanguage := detectLanguage.Language.String()

	targetLanguage := translate.DetermineTargetLanguage(sourceLanguage)

	translatedText, err := translate.TranslateTextWithModel(targetLanguage, text, "nmt")

	if err != nil {

		log.Printf("[DEBUG] Ошибка при переводе текста: %v", err)

		wrongText := "Неправильный текст!"

		wrongText = fmt.Sprintf("Ошибка: \n\n%s", wrongText)

		SendMessageWithContinueAndFinishButton(bot, chatID, wrongText)

		return ""
	}

	log.Printf("[DEBUG] Переведённый текст: %s", translatedText)

	translatedText = fmt.Sprintf("Перевод: \n%s", translatedText)

	SendMessageWithContinueAndFinishButton(bot, chatID, translatedText)

	userStates[chatID] = UserState{TranslationPrompt: true}

	return ""

}
