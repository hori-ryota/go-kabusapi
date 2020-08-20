package kabusapi

import "context"

// APIトークン
type APIToken string

func (c Client) IssuingToken(
	ctx context.Context,
	apiPassword string,
) (IssuingTokenResponse, error) {
	req := NewIssuingTokenRequest(apiPassword)
	res := IssuingTokenResponse{}
	err := c.postRequest(ctx, "/token", req, &res)
	return res, err
}

//genconstructor
type IssuingTokenRequest struct {
	// APIパスワード
	APIPassword string `json:"APIPassword" required:""`
}

type IssuingTokenResponse struct {
	// APIトークン
	Token APIToken `json:"Token"`
}
