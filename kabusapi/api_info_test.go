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

func TestClient_GetSymbolInfo(t *testing.T) {
	expectPathQuery := "/kabusapi/symbol/9433@1"

	respBody := `
{
  "Symbol": "9433",
  "SymbolName": "ＫＤＤＩ",
  "DisplayName": "ＫＤＤＩ",
  "Exchange": 1,
  "ExchangeName": "東証１部",
  "BisCategory": "5250      ",
  "TotalMarketValue": 7654484465100,
  "TotalStocks": 4484,
  "TradingUnit": 100,
  "FiscalYearEndBasic": 20210331,
  "PriceRangeGroup": "10003",
  "KCMarginBuy": true,
  "KCMarginSell": true,
  "MarginBuy": true,
  "MarginSell": true,
  "UpperLimit": 4041,
  "LowerLimit": 2641
}
`

	RESTTestGet(t, true, expectPathQuery, commonHeaderWithAPIKey, respBody,
		func(c kabusapi.Client) {
			res, err := c.GetSymbolInfo(context.Background(), "9433", kabusapi.Tosho)
			assert.NoError(t, err)
			JSONEq(t, respBody, res)
		},
	)
}

func TestClient_GetOrders(t *testing.T) {
	for _, tt := range []struct {
		name            string
		options         []kabusapi.GetOrdersRequestOption
		expectPathQuery string
	}{
		{
			name:            "without options",
			expectPathQuery: "/kabusapi/orders",
		},
		{
			name: "with options",
			options: []kabusapi.GetOrdersRequestOption{
				kabusapi.WithGetOrdersRequestQuery(kabusapi.GetOrdersRequestQueryProductAll),
			},
			expectPathQuery: "/kabusapi/orders?product=0",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			respBody := `
[
	{
		"ID": "20200715A02N04738436",
		"State": 5,
		"OrderState": 5,
		"OrdType": 1,
		"RecvTime": "2020-07-16T18:00:51.763683",
		"Symbol": "8306",
		"SymbolName": "三菱ＵＦＪフィナンシャル・グループ",
		"Exchange": 1,
		"ExchangeName": "東証１部",
		"Price": 704.5,
		"OrderQty": 1500,
		"CumQty": 1500,
		"Side": "1",
		"CashMargin": 2,
		"AccountType": 4,
		"DelivType": 2,
		"ExpireDay": 20200702,
		"MarginTradeType": 1,
		"Details": [
			{
				"SeqNum": 1,
				"ID": "20200715A02N04738436",
				"RecType": 1,
				"ExchangeID": "00000000-0000-0000-0000-00000000",
				"State": 3,
				"TransactTime": "2020-07-16T18:00:51.763683",
				"OrdType": 1,
				"Price": 704.5,
				"Qty": 1500,
				"ExecutionID": "",
				"ExecutionDay": "2020-07-02T18:02:00",
				"DelivDay": 20200706,
				"Commission": 0,
				"CommissionTax": 0
			}
		]
	}
]
`

			RESTTestGet(t, true, tt.expectPathQuery, commonHeaderWithAPIKey, respBody,
				func(c kabusapi.Client) {
					res, err := c.GetOrders(context.Background(), tt.options...)
					assert.NoError(t, err)
					JSONEq(t, respBody, res)
				},
			)
		})
	}

}

func TestClient_GetPositions(t *testing.T) {
	for _, tt := range []struct {
		name            string
		options         []kabusapi.GetPositionsRequestOption
		expectPathQuery string
	}{
		{
			name:            "without options",
			expectPathQuery: "/kabusapi/positions",
		},
		{
			name: "with options",
			options: []kabusapi.GetPositionsRequestOption{
				kabusapi.WithGetPositionsRequestQuery(kabusapi.GetPositionsRequestQueryProductAll),
			},
			expectPathQuery: "/kabusapi/positions?product=0",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			respBody := `
[
  {
    "ExecutionID": "20200715E02N04738464",
    "AccountType": 4,
    "Symbol": "8306",
    "SymbolName": "三菱ＵＦＪフィナンシャル・グループ",
    "Exchange": 1,
    "ExchangeName": "東証１部",
    "ExecutionDay": 20200702,
    "Price": 704,
    "LeavesQty": 500,
    "HoldQty": 0,
    "Side": "1",
    "Expenses": 0,
    "Commission": 1620,
    "CommissionTax": 162,
    "ExpireDay": 20201229,
    "MarginTradeType": 1,
    "CurrentPrice": 414.5,
    "Valuation": 207250,
    "ProfitLoss": 144750,
    "ProfitLossRate": 41.12215909090909
  }
]
`

			RESTTestGet(t, true, tt.expectPathQuery, commonHeaderWithAPIKey, respBody,
				func(c kabusapi.Client) {
					res, err := c.GetPositions(context.Background(), tt.options...)
					assert.NoError(t, err)
					JSONEq(t, respBody, res)
				},
			)
		})
	}

}
