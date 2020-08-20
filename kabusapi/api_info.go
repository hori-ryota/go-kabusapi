package kabusapi

import (
	"context"
	"path"
	"strconv"
	time "time"
)

func (c Client) GetBoard(
	ctx context.Context,
	symbol Symbol,
	exchange Exchange,
) (GetBoardResponse, error) {
	res := GetBoardResponse{}
	err := c.getRequest(ctx, path.Join("/board", string(symbol)+"@"+strconv.Itoa(int(exchange))), nil, &res)
	return res, err
}

type GetBoardResponse struct {
	// 銘柄コード
	Symbol Symbol `json:"Symbol"`
	// 銘柄名
	SymbolName string `json:"SymbolName"`
	// 市場コード
	Exchange Exchange `json:"Exchange"`
	// 市場名称
	ExchangeName string `json:"ExchangeName"`
	// 現値
	CurrentPrice float64 `json:"CurrentPrice"`
	// 現値時刻
	CurrentPriceTime time.Time `json:"CurrentPriceTime"`
	// 現値前値比較
	CurrentPriceChangeStatus CurrentPriceChangeStatus `json:"CurrentPriceChangeStatus"`
	// 現値ステータス
	CurrentPriceStatus CurrentPriceStatus `json:"CurrentPriceStatus"`
	// 計算用現値
	CalcPrice float64 `json:"CalcPrice"`
	// 前日終値
	PreviousClose float64 `json:"PreviousClose"`
	// 前日終値日付
	PreviousCloseTime *time.Time `json:"PreviousCloseTime"`
	// 前日比
	ChangePreviousClose float64 `json:"ChangePreviousClose"`
	// 騰落率
	ChangePreviousClosePer float64 `json:"ChangePreviousClosePer"`
	// 始値
	OpeningPrice float64 `json:"OpeningPrice"`
	// 始値時刻
	OpeningPriceTime time.Time `json:"OpeningPriceTime"`
	// 高値
	HighPrice float64 `json:"HighPrice"`
	// 高値時刻
	HighPriceTime time.Time `json:"HighPriceTime"`
	// 安値
	LowPrice float64 `json:"LowPrice"`
	// 安値時刻
	LowPriceTime time.Time `json:"LowPriceTime"`
	// 売買高
	TradingVolume float64 `json:"TradingVolume"`
	// 売買高時刻
	TradingVolumeTime time.Time `json:"TradingVolumeTime"`
	// 売買高加重平均価格（VWAP）
	VWAP float64 `json:"VWAP"`
	// 売買代金
	TradingValue float64 `json:"TradingValue"`
	// 最良売気配数量
	BidQty float64 `json:"BidQty"`
	// 最良売気配値段
	BidPrice float64 `json:"BidPrice"`
	// 最良売気配時刻
	BidTime time.Time `json:"BidTime"`
	// 最良売気配フラグ
	BidSign QuoteSign `json:"BidSign"`
	// 売成行数量
	MarketOrderSellQty float64 `json:"MarketOrderSellQty"`
	// 売気配数量1本目
	Sell1 Quote `json:"Sell1"`
	// 売気配数量2本目
	Sell2 Quote `json:"Sell2"`
	// 売気配数量3本目
	Sell3 Quote `json:"Sell3"`
	// 売気配数量4本目
	Sell4 Quote `json:"Sell4"`
	// 売気配数量5本目
	Sell5 Quote `json:"Sell5"`
	// 売気配数量6本目
	Sell6 Quote `json:"Sell6"`
	// 売気配数量7本目
	Sell7 Quote `json:"Sell7"`
	// 売気配数量8本目
	Sell8 Quote `json:"Sell8"`
	// 売気配数量9本目
	Sell9 Quote `json:"Sell9"`
	// 売気配数量10本目
	Sell10 Quote `json:"Sell10"`
	// 最良買気配数量
	AskQty float64 `json:"AskQty"`
	// 最良買気配値段
	AskPrice float64 `json:"AskPrice"`
	// 最良買気配時刻
	AskTime time.Time `json:"AskTime"`
	// 最良買気配フラグ
	AskSign QuoteSign `json:"AskSign"`
	// 買成行数量
	MarketOrderBuyQty float64 `json:"MarketOrderBuyQty"`
	// 買気配数量1本目
	Buy1 Quote `json:"Buy1"`
	// 買気配数量2本目
	Buy2 Quote `json:"Buy2"`
	// 買気配数量3本目
	Buy3 Quote `json:"Buy3"`
	// 買気配数量4本目
	Buy4 Quote `json:"Buy4"`
	// 買気配数量5本目
	Buy5 Quote `json:"Buy5"`
	// 買気配数量6本目
	Buy6 Quote `json:"Buy6"`
	// 買気配数量7本目
	Buy7 Quote `json:"Buy7"`
	// 買気配数量8本目
	Buy8 Quote `json:"Buy8"`
	// 買気配数量9本目
	Buy9 Quote `json:"Buy9"`
	// 買気配数量10本目
	Buy10 Quote `json:"Buy10"`
	// Over気配数量
	OverSellQty float64 `json:"OverSellQty"`
	// Under気配数量
	UnderBuyQty float64 `json:"UnderBuyQty"`
	// 時価総額
	TotalMarketValue float64 `json:"TotalMarketValue"`
}

type Quote struct {
	// 時刻
	Time *time.Time `json:"Time,omitempty"`
	// 気配フラグ
	Sign QuoteSign `json:"Sign,omitempty"`
	// 値段
	Price float64 `json:"Price"`
	// 数量
	Qty float64 `json:"Qty"`
}

func (c Client) GetSymbolInfo(
	ctx context.Context,
	symbol Symbol,
	exchange Exchange,
) (GetSymbolInfoResponse, error) {
	res := GetSymbolInfoResponse{}
	err := c.getRequest(ctx, path.Join("/symbol", string(symbol)+"@"+strconv.Itoa(int(exchange))), nil, &res)
	return res, err
}

type GetSymbolInfoResponse struct {
	// 銘柄コード
	Symbol Symbol `json:"Symbol"`
	// 銘柄名
	SymbolName string `json:"SymbolName"`
	// 銘柄略称
	DisplayName string `json:"DisplayName"`
	// 市場コード
	Exchange Exchange `json:"Exchange"`
	// 市場名称
	ExchangeName string `json:"ExchangeName"`
	// 業種コード名
	BisCategory BisCategory `json:"BisCategory"`
}
