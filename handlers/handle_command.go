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
	case "tr":
		return HandleTranslateCommand(update)
	case "start":
		return "Привет я BiB, воспользуйся командой /help и узнай о моих возможностях."
	default:
		return "Default"
	}
}

func HandleTranslateCommand(update tgbotapi.Update) string {
	text := strings.TrimSpace(update.Message.CommandArguments())

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
		return "Произошла ошибка при переводе текста."
	}

	log.Printf("[DEBUG] Переведённый текст: %s", translatedText)

	return fmt.Sprintf("Перевод: \n%s", translatedText)

}
