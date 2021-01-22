package cdc

import (
	"SellAllCro/config"
	"bytes"
	"encoding/json"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
)

const (
	ProdUri = "https://api.crypto.com/v2/"
)

func CdcClient(method string, params map[string]interface{}, reqMethod string) *http.Request {
	req := RequestFormat{
		Id:     888,
		Method: method,
		ApiKey: config.Config.ApiKey,
		Params: params,
		Nonce:  milliTimestamp(),
	}
	signStr := sign(req)
	req.Sig = signStr

	jsonByte, _ := json.Marshal(req)
	r, _ := http.NewRequest(reqMethod, ProdUri+req.Method, bytes.NewBuffer(jsonByte))
	r.Header.Set("Content-Type", "application/json")

	return r
}

func GetCroBalance() (balance decimal.Decimal, err error) {
	var result BalanceResp
	params := map[string]interface{}{
		"currency": "CRO",
	}
	req := CdcClient("private/get-account-summary", params, http.MethodPost)
	jsonByte, err := getResponseJson(req)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonByte, &result)
	if err != nil {
		return
	}

	balance = result.Result.Accounts[0].Available
	log.Printf("Get CRO balance %v", balance)
	return
}

func CreateOrder(croBalance decimal.Decimal) (orderId string, err error) {
	var result CreateOrderResp
	pair := "CRO_" + config.Config.ToCoin
	params := map[string]interface{}{
		"instrument_name": pair,
		"side":            "SELL",
		"type":            "MARKET",
		"quantity":        croBalance,
	}
	req := CdcClient("private/create-order", params, http.MethodPost)
	jsonByte, err := getResponseJson(req)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonByte, &result)
	if err != nil {
		return
	}

	orderId = result.Result.OrderID
	log.Printf("Create order success, order id: %s", orderId)
	return
}

func GetOrderDetail(orderId string) (result OrderDetail, err error) {
	params := map[string]interface{}{
		"order_id": orderId,
	}
	req := CdcClient("private/get-order-detail", params, http.MethodPost)
	jsonByte, err := getResponseJson(req)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonByte, &result)
	if err != nil {
		return
	}

	return
}
