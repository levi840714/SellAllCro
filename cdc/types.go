package cdc

import "github.com/shopspring/decimal"

type (
	RequestFormat struct {
		Id     int64                  `json:"id"`
		Method string                 `json:"method"`
		Params map[string]interface{} `json:"params"`
		ApiKey string                 `json:"api_key"`
		Sig    string                 `json:"sig"`
		Nonce  int64                  `json:"nonce"`
	}

	BalanceResp struct {
		ID     int    `json:"id"`
		Method string `json:"method"`
		Code   int    `json:"code"`
		Result struct {
			Accounts []struct {
				Balance   decimal.Decimal `json:"balance"`
				Available decimal.Decimal `json:"available"`
				Order     float64         `json:"order"`
				Stake     int             `json:"stake"`
				Currency  string          `json:"currency"`
			} `json:"accounts"`
		} `json:"result"`
	}

	CreateOrderResp struct {
		ID     int    `json:"id"`
		Method string `json:"method"`
		Result struct {
			OrderID   string `json:"order_id"`
			ClientOid string `json:"client_oid"`
		} `json:"result"`
	}

	OrderDetail struct {
		ID     int    `json:"id"`
		Method string `json:"method"`
		Code   int    `json:"code"`
		Result struct {
			TradeList []struct {
				Side           string          `json:"side"`
				InstrumentName string          `json:"instrument_name"`
				Fee            float64         `json:"fee"`
				TradeID        string          `json:"trade_id"`
				CreateTime     int64           `json:"create_time"`
				TradedPrice    decimal.Decimal `json:"traded_price"`
				TradedQuantity decimal.Decimal `json:"traded_quantity"`
				FeeCurrency    string          `json:"fee_currency"`
				OrderID        string          `json:"order_id"`
			} `json:"trade_list"`
			OrderInfo struct {
				Status             string          `json:"status"`
				Side               string          `json:"side"`
				OrderID            string          `json:"order_id"`
				ClientOid          string          `json:"client_oid"`
				CreateTime         int64           `json:"create_time"`
				UpdateTime         int64           `json:"update_time"`
				Type               string          `json:"type"`
				InstrumentName     string          `json:"instrument_name"`
				CumulativeQuantity decimal.Decimal `json:"cumulative_quantity"`
				CumulativeValue    decimal.Decimal `json:"cumulative_value"`
				AvgPrice           decimal.Decimal `json:"avg_price"`
				FeeCurrency        string          `json:"fee_currency"`
				TimeInForce        string          `json:"time_in_force"`
				ExecInst           string          `json:"exec_inst"`
			} `json:"order_info"`
		} `json:"result"`
	}
)
