package main

import (
	"log"
	"os"
	botcreateauth "projects/BIb_bot/bot_create_auth"
	"projects/BIb_bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	err error
)

func main() {

	botcreateauth.LoadEnvFile(err)

	token := os.Getenv("CHECK_BOT")
	
	bot := botcreateauth.CreateBot(token)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	userStates := make(map[int64]handlers.UserState)

	for update := range updates {

		handlers.HandleUpdate(update, bot, userStates)
		
	}

}


//После переводенного текста нужно добавить кнопку выход, а если пользователь хочет переводить что то дальше он будет дальше ждать сообщения