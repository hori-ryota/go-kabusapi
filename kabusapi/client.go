package kabusapi

import (
	"bytes"
	"context"
	"encoding/json"
	fmt "fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

const version = "0.0.1"

// Client is implementation of kabuステーションAPI (v1.0) client.
//genconstructor
type Client struct {
	restBaseURL url.URL `required:"url.URL{Scheme: \"http\", Host: \"localhost:18081\", Path: \"kabusapi\"}" setter:"OverwriteRestURLBase"`
	pushBaseURL url.URL `required:"url.URL{Scheme: \"ws\", Host: \"localhost:18081\", Path: \"kabusapi/websocket\"}" setter:"OverwritePushURLBase"`
	apiToken    APIToken
	userAgent   string       `required:"fmt.Sprintf(\"KabusAPIGoClient/%s (%s)\", version, runtime.Version())" setter:""`
	httpClient  *http.Client `required:"http.DefaultClient" setter:""`
}

// NewTestingClient is constructor of Client for testing
func NewTestingClient() Client {
	c := NewClient()
	c.OverwriteRestURLBase(url.URL{Scheme: "http", Host: "localhost:18082", Path: "kabusapi"})
	c.OverwriteRestURLBase(url.URL{Scheme: "ws", Host: "localhost:18082", Path: "kabusapi/websocket"})
	return c
}

func (c *Client) Initialize(
	ctx context.Context,
	apiPassword string,
) error {
	res, err := c.IssuingToken(ctx, apiPassword)
	if err != nil {
		return fmt.Errorf("failed to IssuingToken: %w", err)
	}
	c.apiToken = res.Token
	return nil
}

func (c Client) newRESTRequest(ctx context.Context, method string, requestPath string, query url.Values, body io.Reader) (*http.Request, error) {
	u := c.restBaseURL
	u.Path = path.Join(u.Path, requestPath)
	u.RawQuery = query.Encode()
	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("X-API-KEY", string(c.apiToken))

	return req, nil
}

func (c Client) executeRESTRequest(req *http.Request, res interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		e := ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return fmt.Errorf("failed to request: failed to decode requestError: %w", err)
		}
		e.HTTPStatus = resp.Status
		e.HTTPStatusCode = resp.StatusCode
		return e
	}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	return nil
}

func (c Client) restRequestWithoutParams(ctx context.Context, method string, requestPath string, res interface{}) error {
	req, err := c.newRESTRequest(ctx, method, requestPath, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to generate REST request: %w", err)
	}
	return c.executeRESTRequest(req, res)
}

func (c Client) restRequestWithBody(ctx context.Context, method string, requestPath string, body interface{}, res interface{}) error {
	if body == nil {
		return c.restRequestWithoutParams(ctx, method, requestPath, res)
	}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(body); err != nil {
		return fmt.Errorf("failed to encode body to json: %w", err)
	}
	req, err := c.newRESTRequest(ctx, method, requestPath, nil, b)
	if err != nil {
		return fmt.Errorf("failed to generate REST request: %w", err)
	}
	return c.executeRESTRequest(req, res)
}

func (c Client) postRequest(ctx context.Context, requestPath string, body interface{}, res interface{}) error {
	return c.restRequestWithBody(ctx, http.MethodPost, requestPath, body, res)
}
func (c Client) putRequest(ctx context.Context, requestPath string, body interface{}, res interface{}) error {
	return c.restRequestWithBody(ctx, http.MethodPut, requestPath, body, res)
}

func (c Client) getRequest(ctx context.Context, requestPath string, query url.Values, res interface{}) error {
	req, err := c.newRESTRequest(ctx, http.MethodGet, requestPath, query, nil)
	if err != nil {
		return fmt.Errorf("failed to generate REST request: %w", err)
	}
	return c.executeRESTRequest(req, res)
}
