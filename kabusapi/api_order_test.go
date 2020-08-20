package kabusapi_test

import (
	"context"
	"testing"

	"github.com/hori-ryota/go-kabusapi/kabusapi"
	"github.com/stretchr/testify/assert"
)

func TestClient_SendOrder(t *testing.T) {
	expectRequestBody := `
{
  "Password": "xxxxxx",
  "Symbol": "3997",
  "Exchange": 1,
  "SecurityType": 1,
  "Side": "2",
  "CashMargin": 1,
  "MarginTradeType": 1,
  "DelivType": 2,
  "FundType": "AA",
  "AccountType": 2,
  "Qty": 100,
  "ClosePositionOrder": null,
  "ClosePositions": [
    {
      "HoldID": "20200715E02N04738464",
      "Qty": 100
    }
  ],
  "Price": 2000,
  "ExpireDay": 20200716,
  "FrontOrderType": 20
}
`
	respBody := `
{
  "Result": 0,
  "OrderId": "20200529A01N06848002"
}
`

	RESTTestPost(t, true, "/kabusapi/sendorder", commonHeaderWithAPIKey, expectRequestBody, respBody,
		func(c kabusapi.Client) {
			res, err := c.SendOrder(context.Background(), kabusapi.SendOrderRequest{
				Password:        "xxxxxx",
				Symbol:          "3997",
				Exchange:        kabusapi.Tosho,
				SecurityType:    kabusapi.Stock,
				Side:            kabusapi.Buy,
				CashMargin:      kabusapi.Genbutsu,
				MarginTradeType: kabusapi.SystemMarginTrade,
				DelivType:       kabusapi.Deposit,
				FundType:        kabusapi.ShinyoSubstitute,
				AccountType:     kabusapi.GeneralAccount,
				Qty:             100,
				ClosePositions: []kabusapi.ClosePosition{
					{
						HoldID: "20200715E02N04738464",
						Qty:    100,
					},
				},
				Price:          2000,
				ExpireDay:      kabusapi.NewDate(2020, 7, 16),
				FrontOrderType: kabusapi.Sashine,
			})
			assert.NoError(t, err)
			JSONEqWithoutResultCode(t, respBody, res)
		},
	)
}

func TestClient_CancelOrder(t *testing.T) {
	expectRequestBody := `
{
  "OrderId": "20200529A01N06848002",
  "Password": "xxxxxx"
}
`
	respBody := `
{
  "Result": 0,
  "OrderId": "20200529A01N06848002"
}
`

	RESTTestPut(t, true, "/kabusapi/cancelorder", commonHeaderWithAPIKey, expectRequestBody, respBody,
		func(c kabusapi.Client) {
			res, err := c.CancelOrder(context.Background(), "20200529A01N06848002", "xxxxxx")
			assert.NoError(t, err)
			JSONEqWithoutResultCode(t, respBody, res)
		},
	)
}
