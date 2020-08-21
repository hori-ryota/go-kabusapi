package kabusapi_test

import (
	"context"
	"testing"

	"github.com/hori-ryota/go-kabusapi/kabusapi"
	"github.com/stretchr/testify/assert"
)

func TestClient_RegisterSymbols(t *testing.T) {
	expectRequestBody := `
{
  "Symbols": [
    {
      "Symbol": "9433",
      "Exchange": 1
    }
  ]
}
`
	respBody := `
{
  "RegistList": [
    {
      "Symbol": "9433",
      "Exchange": 1
    }
  ]
}
`

	RESTTestPut(t, true, "/kabusapi/register", commonHeaderWithAPIKey, expectRequestBody, respBody,
		func(c kabusapi.Client) {
			res, err := c.RegisterSymbols(context.Background(), kabusapi.NewSymbolItem("9433", kabusapi.Tosho))
			assert.NoError(t, err)
			JSONEqWithoutResultCode(t, respBody, res)
		},
	)
}

func TestClient_UnregisterSymbols(t *testing.T) {
	expectRequestBody := `
{
  "Symbols": [
    {
      "Symbol": "9433",
      "Exchange": 1
    }
  ]
}
`
	respBody := `
{
  "RegistList": [
    {
      "Symbol": "9433",
      "Exchange": 1
    }
  ]
}
`

	RESTTestPut(t, true, "/kabusapi/unregister", commonHeaderWithAPIKey, expectRequestBody, respBody,
		func(c kabusapi.Client) {
			res, err := c.UnregisterSymbols(context.Background(), kabusapi.NewSymbolItem("9433", kabusapi.Tosho))
			assert.NoError(t, err)
			JSONEqWithoutResultCode(t, respBody, res)
		},
	)
}

func TestClient_UnregisterAllSymbols(t *testing.T) {

	RESTTestPut(t, true, "/kabusapi/unregister/all", commonHeaderWithAPIKey, "", "",
		func(c kabusapi.Client) {
			err := c.UnregisterAllSymbols(context.Background())
			assert.NoError(t, err)
		},
	)
}
