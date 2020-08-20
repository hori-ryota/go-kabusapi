package kabusapi_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/hori-ryota/go-kabusapi/kabusapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func RESTTestGet(
	t *testing.T,
	needsInitializeClient bool,
	expectPathQuery string,
	requiredHeader map[string]string,
	responseJSON string,
	doRequest func(c kabusapi.Client),
) {
	t.Helper()
	RESTTest(
		t,
		needsInitializeClient,
		http.MethodGet,
		expectPathQuery,
		requiredHeader, "", responseJSON, doRequest,
	)
}

func RESTTestPost(
	t *testing.T,
	needsInitializeClient bool,
	expectPathQuery string,
	requiredHeader map[string]string,
	expectRequestJSON string,
	responseJSON string,
	doRequest func(c kabusapi.Client),
) {
	t.Helper()
	RESTTest(
		t,
		needsInitializeClient,
		http.MethodPost,
		expectPathQuery,
		requiredHeader,
		expectRequestJSON,
		responseJSON,
		doRequest,
	)
}

func RESTTestPut(
	t *testing.T,
	needsInitializeClient bool,
	expectPathQuery string,
	requiredHeader map[string]string,
	expectRequestJSON string,
	responseJSON string,
	doRequest func(c kabusapi.Client),
) {
	t.Helper()
	RESTTest(
		t,
		needsInitializeClient,
		http.MethodPut,
		expectPathQuery,
		requiredHeader,
		expectRequestJSON,
		responseJSON,
		doRequest,
	)
}

func RESTTest(
	t *testing.T,
	needsInitializeClient bool,
	expectMethod string,
	expectPathQuery string,
	requiredHeader map[string]string,
	expectRequestJSON string,
	responseJSON string,
	doRequest func(c kabusapi.Client),
) {
	t.Helper()
	ts := RESTTestServer(
		t,
		expectMethod,
		expectPathQuery,
		requiredHeader,
		expectRequestJSON,
		responseJSON,
	)
	defer ts.Close()
	client := kabusapi.NewClient()

	if needsInitializeClient {
		InitializeClient(t, &client)
	}

	u, err := url.Parse(ts.URL)
	require.NoError(t, err)
	u.Path = "/kabusapi"
	client.OverwriteRestURLBase(*u)
	doRequest(client)
}

func RESTTestServer(
	t *testing.T,
	expectMethod string,
	expectPathQuery string,
	requiredHeader map[string]string,
	expectRequestJSON string,
	responseJSON string,
) *httptest.Server {
	t.Helper()
	u, err := url.Parse(expectPathQuery)
	require.NoError(t, err)
	u.RawQuery = u.Query().Encode()
	expectPath := u.Path
	expectQuery := u.Query()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, expectMethod, req.Method)
		assert.Equal(t, expectPath, req.URL.Path)
		assert.Equal(t, expectQuery, req.URL.Query())
		for k, v := range requiredHeader {
			assert.Equal(t, v, req.Header.Get(k))
		}
		if expectRequestJSON != "" {
			b, err := ioutil.ReadAll(req.Body)
			assert.NoError(t, err)
			JSONEq(t, expectRequestJSON, b)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(responseJSON))
		require.NoError(t, err)
	}))
}

func JSONEqWithoutResultCode(t *testing.T, expect string, actual interface{}) {
	t.Helper()
	expect = strings.ReplaceAll(expect, `"Result": 0,`, "")
	JSONEq(t, expect, actual)
}

func JSONEq(t *testing.T, expect string, actual interface{}) {
	t.Helper()
	if s, ok := actual.(string); ok {
		assert.JSONEq(t, expect, s)
		return
	}
	if s, ok := actual.([]byte); ok {
		assert.JSONEq(t, expect, string(s))
		return
	}
	j, err := json.Marshal(actual)
	assert.NoError(t, err)
	assert.JSONEq(t, expect, string(j))
}

func InitializeClient(
	t *testing.T,
	client *kabusapi.Client,
) {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`
{
	"Result": 0,
	"Token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
`))
		require.NoError(t, err)
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	require.NoError(t, err)
	client.OverwriteRestURLBase(*u)

	client.Initialize(context.Background(), "apiPassword")
}
