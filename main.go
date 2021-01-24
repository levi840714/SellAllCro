package main

import (
	"SellAllCro/cdc"
	"SellAllCro/config"
	"SellAllCro/telegram"
	"log"
	"time"
)

func init() {
	err := config.SetConfigFile("config.json")
	if err != nil {
		log.Fatalf("set config failed, err: %s", err)
	}

	// initial telegram bot, if you want notify to telegram
	if config.Config.TgBotToken != "" && config.Config.TgChannelID != 0 {
		telegram.Init()
	}
}

func main() {
	log.Println("Start sell all CRO automatically on crypto.com exchange")

	for {
		// init websocket connect and subscribe order channel
		ws := cdc.NewWebsocket()
		ws.InitWebsocket()

		balance, err := cdc.GetCroBalance()
		if err != nil {
			log.Printf("Get account CRO balance failed, err: %s", err)
			continue
		}

		orderId, err := cdc.CreateOrder(balance)
		if err != nil && orderId == "" {
			log.Printf("Create sell CRO order failed, err: %s", err)
		}

		ws.Close()
		log.Println("end selling")

		time.Sleep(time.Hour * 1)
	}
}
