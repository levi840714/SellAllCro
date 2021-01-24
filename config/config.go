package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Config *AppConfig = nil

type AppConfig struct {
	ApiKey      string `json:"api_key"`
	SecretKey   string `json:"secret_key"`
	ToCoin      string `json:"to_coin"`
	TgBotToken  string `json:"tg_bot_token"`
	TgChannelID int64  `json:"tg_channel_id"`
}

func SetConfigFile(file string) (err error) {
	var conf AppConfig
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	r, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	err = json.Unmarshal(r, &conf)
	if err != nil {
		return
	}

	Config = &conf
	return
}
