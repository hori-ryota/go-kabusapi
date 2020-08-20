package kabusapi

import (
	"context"
	"path"
	"strconv"
)

func (c Client) GetWalletCash(
	ctx context.Context,
) (GetWalletCashResponse, error) {
	res := GetWalletCashResponse{}
	err := c.getRequest(ctx, "/wallet/cash", nil, &res)
	return res, err
}

type GetWalletCashResponse struct {
	// 現物買付可能額
	StockAccountWallet float64 `json:"StockAccountWallet"`
}

func (c Client) GetWalletCashWithSymbol(
	ctx context.Context,
	symbol Symbol,
	exchange Exchange,
) (GetWalletCashResponse, error) {
	res := GetWalletCashResponse{}
	err := c.getRequest(ctx, path.Join("/wallet/cash", string(symbol)+"@"+strconv.Itoa(int(exchange))), nil, &res)
	return res, err
}

func (c Client) GetWalletMargin(
	ctx context.Context,
	symbol Symbol,
	exchange Exchange,
) (GetWalletMarginResponse, error) {
	res := GetWalletMarginResponse{}
	err := c.getRequest(ctx, path.Join("/wallet/margin", string(symbol)+"@"+strconv.Itoa(int(exchange))), nil, &res)
	return res, err
}

type GetWalletMarginResponse struct {
	// 信用新規可能額
	MarginAccountWallet float64 `json:"MarginAccountWallet"`
	// 保証金維持率
	DepositkeepRate float64 `json:"DepositkeepRate"`
	// 委託保証金率
	ConsignmentDepositRate float64 `json:"ConsignmentDepositRate"`
	// 現金委託保証金率
	CashOfConsignmentDepositRate float64 `json:"CashOfConsignmentDepositRate"`
}
