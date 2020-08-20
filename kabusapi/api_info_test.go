package kabusapi_test

import (
	"context"
	"testing"

	"github.com/hori-ryota/go-kabusapi/kabusapi"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetBoard(t *testing.T) {
	expectPathQuery := "/kabusapi/board/5401@1"

	respBody := `
{
  "Symbol": "5401",
  "SymbolName": "新日鐵住金",
  "Exchange": 1,
  "ExchangeName": "東証１部",
  "CurrentPrice": 2408,
  "CurrentPriceTime": "2020-07-22T15:00:00+09:00",
  "CurrentPriceChangeStatus": "0058",
  "CurrentPriceStatus": 1,
  "CalcPrice": 343.7,
  "PreviousClose": 1048,
  "PreviousCloseTime": null,
  "ChangePreviousClose": 1360,
  "ChangePreviousClosePer": 129.77,
  "OpeningPrice": 2380,
  "OpeningPriceTime": "2020-07-22T09:00:00+09:00",
  "HighPrice": 2418,
  "HighPriceTime": "2020-07-22T13:25:47+09:00",
  "LowPrice": 2370,
  "LowPriceTime": "2020-07-22T10:00:04+09:00",
  "TradingVolume": 4571500,
  "TradingVolumeTime": "2020-07-22T15:00:00+09:00",
  "VWAP": 2394.4262,
  "TradingValue": 10946119350,
  "BidQty": 100,
  "BidPrice": 2408.5,
  "BidTime": "2020-07-22T14:59:59+09:00",
  "BidSign": "0101",
  "MarketOrderSellQty": 0,
  "Sell1": {
    "Time": "2020-07-22T14:59:59+09:00",
    "Sign": "0101",
    "Price": 2408.5,
    "Qty": 100
  },
  "Sell2": {
    "Price": 2409,
    "Qty": 800
  },
  "Sell3": {
    "Price": 2409.5,
    "Qty": 2100
  },
  "Sell4": {
    "Price": 2410,
    "Qty": 800
  },
  "Sell5": {
    "Price": 2410.5,
    "Qty": 500
  },
  "Sell6": {
    "Price": 2411,
    "Qty": 8400
  },
  "Sell7": {
    "Price": 2411.5,
    "Qty": 1200
  },
  "Sell8": {
    "Price": 2412,
    "Qty": 27200
  },
  "Sell9": {
    "Price": 2412.5,
    "Qty": 400
  },
  "Sell10": {
    "Price": 2413,
    "Qty": 16400
  },
  "AskQty": 200,
  "AskPrice": 2407.5,
  "AskTime": "2020-07-22T14:59:59+09:00",
  "AskSign": "0101",
  "MarketOrderBuyQty": 0,
  "Buy1": {
    "Time": "2020-07-22T14:59:59+09:00",
    "Sign": "0101",
    "Price": 2407.5,
    "Qty": 200
  },
  "Buy2": {
    "Price": 2407,
    "Qty": 400
  },
  "Buy3": {
    "Price": 2406.5,
    "Qty": 1000
  },
  "Buy4": {
    "Price": 2406,
    "Qty": 5800
  },
  "Buy5": {
    "Price": 2405.5,
    "Qty": 7500
  },
  "Buy6": {
    "Price": 2405,
    "Qty": 2200
  },
  "Buy7": {
    "Price": 2404.5,
    "Qty": 16700
  },
  "Buy8": {
    "Price": 2404,
    "Qty": 30100
  },
  "Buy9": {
    "Price": 2403.5,
    "Qty": 1300
  },
  "Buy10": {
    "Price": 2403,
    "Qty": 3000
  },
  "OverSellQty": 974900,
  "UnderBuyQty": 756000,
  "TotalMarketValue": 3266254659361.4
}
`

	RESTTestGet(t, true, expectPathQuery, commonHeaderWithAPIKey, respBody,
		func(c kabusapi.Client) {
			res, err := c.GetBoard(context.Background(), "5401", kabusapi.Tosho)
			assert.NoError(t, err)
			JSONEq(t, respBody, res)
		},
	)
}
