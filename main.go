package main

import (
	"SellAllCro/config"
	"fmt"
	"log"
)

const (
	ProdUri = "https://api.crypto.com/v2/"
)

func init() {
	err := config.SetConfigFile("config.json")
	if err != nil {
		log.Fatalf("set config failed, err: %s", err)
	}
}

func main() {
	fmt.Println("SellAllCro!!!")
}
