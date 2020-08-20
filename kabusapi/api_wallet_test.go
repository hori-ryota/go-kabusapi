package kabusapi_test

import (
	"context"
	"testing"

	"github.com/hori-ryota/go-kabusapi/kabusapi"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetWalletCash(t *testing.T) {
	expectPathQuery := "/kabusapi/wallet/cash"

	respBody := `
{
  "StockAccountWallet": 0
}
`

	RESTTestGet(t, true, expectPathQuery, commonHeaderWithAPIKey, respBody,
		func(c kabusapi.Client) {
			res, err := c.GetWalletCash(context.Background())
			assert.NoError(t, err)
			JSONEq(t, respBody, res)
		},
	)
}

func TestClient_GetWalletCashWithSymbol(t *testing.T) {
	expectPathQuery := "/kabusapi/wallet/cash/0000@1"

	respBody := `
{
  "StockAccountWallet": 0
}
`

	RESTTestGet(t, true, expectPathQuery, commonHeaderWithAPIKey, respBody,
		func(c kabusapi.Client) {
			res, err := c.GetWalletCashWithSymbol(context.Background(), "0000", kabusapi.Tosho)
			assert.NoError(t, err)
			JSONEq(t, respBody, res)
		},
	)
}

func TestClient_GetWalletMargin(t *testing.T) {
	expectPathQuery := "/kabusapi/wallet/margin/0000@1"

	respBody := `
{
  "MarginAccountWallet": 0,
  "DepositkeepRate": 0,
  "ConsignmentDepositRate": 0,
  "CashOfConsignmentDepositRate": 0
}
`

	RESTTestGet(t, true, expectPathQuery, commonHeaderWithAPIKey, respBody,
		func(c kabusapi.Client) {
			res, err := c.GetWalletMargin(context.Background(), "0000", kabusapi.Tosho)
			assert.NoError(t, err)
			JSONEq(t, respBody, res)
		},
	)
}
