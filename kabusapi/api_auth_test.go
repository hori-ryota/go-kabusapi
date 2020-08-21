package kabusapi_test

import (
	"context"
	"testing"

	"github.com/hori-ryota/go-kabusapi/kabusapi"
	"github.com/stretchr/testify/assert"
)

func TestClient_IssuingToken(t *testing.T) {
	expectRequestBody := `
{
	"APIPassword": "xxxxxx"
}
`
	respBody := `
{
	"ResultCode": 0,
	"Token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
`

	RESTTestPost(t, false, "/kabusapi/token", commonHeader, expectRequestBody, respBody,
		func(c kabusapi.Client) {
			res, err := c.IssuingToken(context.Background(), "xxxxxx")
			assert.NoError(t, err)
			JSONEqWithoutResultCode(t, respBody, res)
		},
	)
}
