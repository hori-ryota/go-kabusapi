package kabusapi

import "context"

//genconstructor
type SymbolItem struct {
	// 銘柄コード
	Symbol Symbol `json:"Symbol" required:""`
	// 事情コード
	Exchange Exchange `json:"Exchange" required:""`
}

func (c Client) RegisterSymbols(
	ctx context.Context,
	symbols ...SymbolItem,
) (RegisterSymbolsResponse, error) {
	req := NewRegisterSymbolsRequest(symbols)
	res := RegisterSymbolsResponse{}
	err := c.putRequest(ctx, "/register", req, &res)
	return res, err
}

//genconstructor
type RegisterSymbolsRequest struct {
	// 登録する銘柄のリスト
	Symbols []SymbolItem `json:"Symbols" required:""`
}

type RegisterSymbolsResponse struct {
	// 現在登録されている銘柄のリスト
	RegisterList []SymbolItem `json:"RegistList"`
}

func (c Client) UnregisterSymbols(
	ctx context.Context,
	symbols ...SymbolItem,
) (UnregisterSymbolsResponse, error) {
	req := NewUnregisterSymbolsRequest(symbols)
	res := UnregisterSymbolsResponse{}
	err := c.putRequest(ctx, "/unregister", req, &res)
	return res, err
}

//genconstructor
type UnregisterSymbolsRequest struct {
	// 登録解除する銘柄のリスト
	Symbols []SymbolItem `json:"Symbols" required:""`
}

type UnregisterSymbolsResponse struct {
	// 現在登録されている銘柄のリスト
	UnregisterList []SymbolItem `json:"RegistList"`
}

func (c Client) UnregisterAllSymbols(
	ctx context.Context,
) error {
	return c.putRequest(ctx, "/unregister/all", nil, nil)
}
