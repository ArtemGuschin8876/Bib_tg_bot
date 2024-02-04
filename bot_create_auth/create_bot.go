package botcreateauth

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var ()

func LoadEnvFile(err error) {

	err = godotenv.Load()

	if err != nil {
		log.Fatal("[ERR] Error loading .env file")
	}
}

func CreateBot(token string) *tgbotapi.BotAPI {

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal("[ERR] The bot has not been created")
	}

	bot.Debug = true

	return bot
}
