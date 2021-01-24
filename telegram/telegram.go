package telegram

import (
	"SellAllCro/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI = nil

func Init() {
	var err error
	bot, err = tgbotapi.NewBotAPI(config.Config.TgBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized telegram on account %s", bot.Self.UserName)
}

func Send(message string) {
	if config.Config.TgBotToken != "" && config.Config.TgChannelID != 0 {
		channel := config.Config.TgChannelID
		msg := tgbotapi.NewMessage(channel, message)

		bot.Send(msg)
	}
}

func OnMessage() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	}
}
