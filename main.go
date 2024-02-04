package main

import (
	"fmt"
	"log"
	"os"
	botcreateauth "projects/BIb_bot/bot_create_auth"
	"projects/BIb_bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	msgText string
	err     error
)

func main() {
	botcreateauth.LoadEnvFile(err)

	token := os.Getenv("CHECK_BOT")

	bot := botcreateauth.CreateBot(token)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		// userID := update.Message.Chat.ID

		handlers.HandleUpdate(update, bot)

		if update.Message.IsCommand() {

			handlers.HandleCommand(update)

		}
		fmt.Println("4")
		// handlers.SendMessage(userID, msgText, bot)
		fmt.Println("5")
	}

}
