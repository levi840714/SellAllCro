package cdc

import (
	"SellAllCro/config"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

const (
	wsUserAPi    = "wss://stream.crypto.com/v2/user"
	wsMarketData = "wss://stream.crypto.com/v2/market"
)

type Action func([]byte)

type WsClient struct {
	Conn   *websocket.Conn
	Method map[string]Action
}

func NewWebsocket() *WsClient {
	ws := &WsClient{}
	ws.Method = map[string]Action{
		"public/heartbeat": ws.RespondHeartbeat,
		"subscribe":        ws.ReceiveOrderChannel,
	}
	return ws
}

func (ws *WsClient) InitWebsocket() {
	ws.Dial()
	go ws.ListenMessage()
	ws.AuthWebsocket()
	ws.SubscribeOrder()
}

func (ws *WsClient) Dial() {
	c, _, err := websocket.DefaultDialer.Dial(wsUserAPi, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	ws.Conn = c
}

func (ws *WsClient) ListenMessage() {
	for {
		var result interface{}

		_, msg, err := ws.Conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("msg: ", string(msg))
		json.Unmarshal(msg, &result)

		resultMap := result.(map[string]interface{})
		method := resultMap["method"].(string)
		if actionFunc, ok := ws.Method[method]; ok {
			actionFunc(msg)
		}
	}
}

func (ws *WsClient) AuthWebsocket() {
	req := RequestFormat{
		Id:     888,
		Method: "public/auth",
		ApiKey: config.Config.ApiKey,
		Params: nil,
		Nonce:  milliTimestamp(),
	}
	signStr := sign(req)
	req.Sig = signStr

	jsonByte, _ := json.Marshal(req)
	ws.Conn.WriteMessage(websocket.TextMessage, jsonByte)
}

func (ws *WsClient) RespondHeartbeat(response []byte) {
	var heartbeat Heartbeat

	json.Unmarshal(response, &heartbeat)
	reply := ReplyHeartbeat{
		Id:     heartbeat.Id,
		Method: "public/respond-heartbeat",
	}

	jsonByte, _ := json.Marshal(reply)
	ws.Conn.WriteMessage(websocket.TextMessage, jsonByte)
}

func (ws *WsClient) SubscribeOrder() {
	params := map[string]interface{}{
		"channels": []string{"user.order.CRO_" + config.Config.ToCoin},
	}
	req := RequestFormat{
		Id:     888,
		Method: "subscribe",
		Params: params,
		Nonce:  milliTimestamp(),
	}

	jsonByte, _ := json.Marshal(req)
	ws.Conn.WriteMessage(websocket.TextMessage, jsonByte)
}

func (ws *WsClient) ReceiveOrderChannel(response []byte) {
	var order SubscribeChannel
	json.Unmarshal(response, &order)

	orderData := order.Result.Data
	for _, data := range orderData {
		if data.Status == "FILLED" {
			log.Printf("Sell %v CRO to %v %s success!!", data.CumulativeQuantity, data.CumulativeValue, config.Config.ToCoin)
		}
	}
}
