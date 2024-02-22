package main

import (
	"log"
	"os"
	"os/signal"
	botcreateauth "projects/BIb_bot/bot_create_auth"
	"projects/BIb_bot/handlers"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	updateID           = 0
	updateTimeout      = 60
	environmentVarName = "TOKEN_BOT"
)

// Создание бота, апдейты, завершение работы бота
func main() {

	err := botcreateauth.LoadEnvFile()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	token := os.Getenv(environmentVarName)

	bot := botcreateauth.CreateBot(token)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(updateID)
	u.Timeout = updateTimeout
	updates := bot.GetUpdatesChan(u)

	userStates := make(map[int64]handlers.UserState)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stopChan
		log.Println("Received interrupt signal. Shutting down gracefully...")
		os.Exit(0)
	}()

	for update := range updates {

		handlers.HandleUpdate(update, bot, userStates)

	}

}
