package main

import (
	"SellAllCro/cdc"
	"SellAllCro/config"
	"fmt"
	"log"
	"time"
)

func init() {
	err := config.SetConfigFile("config.json")
	if err != nil {
		log.Fatalf("set config failed, err: %s", err)
	}
}

func main() {
	log.Println("Start sell all CRO automatically on crypto.com exchange")

	for {
		balance, err := cdc.GetCroBalance()
		if err != nil {
			log.Printf("Get account CRO balance failed, err: %s", err)
			continue
		}

		orderId, err := cdc.CreateOrder(balance)
		if err != nil {
			log.Printf("Create sell CRO order failed, err: %s", err)
		}

		order, err := cdc.GetOrderDetail(orderId)
		if err != nil {
			log.Printf("Get order detail failed, err: %s", err)
		}

		if order.Code == 0 && order.Result.OrderInfo.Status == "FILLED" {
			fmt.Printf("Sell %v CRO to %v %s success!!", order.Result.OrderInfo.CumulativeQuantity, order.Result.OrderInfo.CumulativeValue, config.Config.ToCoin)
		}

		time.Sleep(time.Hour * 1)
	}
}
