package main

import (
	"fmt"
	"log"
	"os"
	"projects/BIb_bot/translate"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (

	msgText string
)
func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERR] Error loading .env file")
	}

	token := os.Getenv("TOKEN_BOT")
	

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("[ERR] The bot has not been created")
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand(){
		switch update.Message.Command(){
		case "list":
			msgText = "Здесь весь лист"
		case "help":
			msgText = "Тут помощь"	
		case "tr":
			text := strings.TrimSpace(update.Message.CommandArguments())
				if text == "" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, введите текст для перевода после команды /tr.")
					bot.Send(msg)
					continue
				}

				log.Printf("[DEBUG] Переводим текст: %s", text)

				translatedText, err := translate.TranslateText(text, "en", "ru")
				if err != nil {
					log.Printf("Error translating text: %v", err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка при переводе текста.")
					bot.Send(msg)
				} else {
					log.Printf("[DEBUG] Получен переведенный текст: %s", translatedText)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Переведенный текст:\n%s", translatedText))
					bot.Send(msg)
				}
		default :
			msgText = "Nothing..."
		}
	}
		userID := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(userID, msgText)

		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
	
}


