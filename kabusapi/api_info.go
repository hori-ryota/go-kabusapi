package kabusapi

import (
	"context"
	url "net/url"
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
	// 時価総額
	TotalMarketValue float64 `json:"TotalMarketValue"`
	// 発行済み株式数（千株）
	TotalStocks float64 `json:"TotalStocks"`
	// 売買単位
	TradingUnit float64 `json:"TradingUnit"`
	// 決算期日
	FiscalYearEndBasic Date `json:"FiscalYearEndBasic"`
	// 呼値グループ
	PriceRangeGroup PriceRangeCode `json:"PriceRangeGroup"`
	// 一般信用買建フラグ
	KCMarginBuy bool `json:"KCMarginBuy"`
	// 一般信用売建フラグ
	KCMarginSell bool `json:"KCMarginSell"`
	// 制度信用買建フラグ
	MarginBuy bool `json:"MarginBuy"`
	// 制度信用売建フラグ
	MarginSell bool `json:"MarginSell"`
	// 値幅上限
	UpperLimit float64 `json:"UpperLimit"`
	// 値幅下限
	LowerLimit float64 `json:"LowerLimit"`
}

func (c Client) GetOrders(
	ctx context.Context,
	options ...GetOrdersRequestOption,
) ([]GetOrdersResponse, error) {
	o := getOrdersRequestOptions{}
	for _, opt := range options {
		opt(&o)
	}

	values := url.Values{}
	if o.Product != nil {
		values.Set("product", string(*o.Product))
	}
	res := []GetOrdersResponse{}
	err := c.getRequest(ctx, "/orders", values, &res)
	return res, err
}

type GetOrdersRequestOption func(*getOrdersRequestOptions)

type GetOrdersRequestQueryProduct string

const (
	GetOrdersRequestQueryProductAll      GetOrdersRequestQueryProduct = "0"
	GetOrdersRequestQueryProductGenbutsu GetOrdersRequestQueryProduct = "1"
	GetOrdersRequestQueryProductMargin   GetOrdersRequestQueryProduct = "2"
)

type getOrdersRequestOptions struct {
	Product *GetOrdersRequestQueryProduct
}

func WithGetOrdersRequestQuery(q GetOrdersRequestQueryProduct) GetOrdersRequestOption {
	return func(o *getOrdersRequestOptions) {
		o.Product = &q
	}
}

type GetOrdersResponse struct {
	// 注文番号
	ID OrderID `json:"ID"`
	// 状態
	State OrderState `json:"State"`
	// 注文状態
	OrderState OrderState `json:"OrderState"`
	// 執行条件
	OrdType OrderType `json:"OrdType"`
	// 受注日次
	RecvTime DateTime `json:"RecvTime"`
	// 銘柄コード
	Symbol Symbol `json:"Symbol"`
	// 銘柄名
	SymbolName string `json:"SymbolName"`
	// 市場コード
	Exchange Exchange `json:"Exchange"`
	// 市場名
	ExchangeName string `json:"ExchangeName"`
	// 値段
	Price float64 `json:"Price"`
	// 発注数量
	OrderQty float64 `json:"OrderQty"`
	// 約定数量
	CumQty float64 `json:"CumQty"`
	// 売買区分
	Side Side `json:"Side"`
	// 現物信用区分
	CashMargin CashMargin `json:"CashMargin"`
	// 口座種別
	AccountType AccountType `json:"AccountType"`
	// 受渡区分
	DelivType DelivType `json:"DelivType"`
	// 注文有効期限
	ExpireDay Date `json:"ExpireDay"`
	// 信用取引区分
	MarginTradeType MarginTradeType `json:"MarginTradeType"`
	// 注文明細
	Details []OrderDetail `json:"Details"`
}

type OrderDetail struct {
	// 通番
	SeqNum int32 `json:"SeqNum"`
	// 注文詳細番号
	ID string `json:"ID"`
	// 明細種別
	RecType RecType `json:"RecType"`
	// 取引所番号
	ExchangeID string `json:"ExchangeID"`
	// 状態
	State OrderDetailState `json:"State"`
	// 処理時刻
	TransactTime DateTime `json:"TransactTime"`
	// 執行条件
	OrdType OrderType `json:"OrdType"`
	// 値段
	Price float64 `json:"Price"`
	// 数量
	Qty float64 `json:"Qty"`
	// 約定番号
	ExecutionID string `json:"ExecutionID"`
	// 約定日次
	ExecutionDay DateTime `json:"ExecutionDay"`
	// 受渡日
	DelivDay Date `json:"DelivDay"`
	// 手数料
	Commission float64 `json:"Commission"`
	// 手数料消費税
	CommissionTax float64 `json:"CommissionTax"`
}

func (c Client) GetPositions(
	ctx context.Context,
	options ...GetPositionsRequestOption,
) ([]GetPositionsResponse, error) {
	o := getPositionsRequestOptions{}
	for _, opt := range options {
		opt(&o)
	}

	values := url.Values{}
	if o.Product != nil {
		values.Set("product", string(*o.Product))
	}
	res := []GetPositionsResponse{}
	err := c.getRequest(ctx, "/positions", values, &res)
	return res, err
}

type GetPositionsRequestOption func(*getPositionsRequestOptions)

type GetPositionsRequestQueryProduct string

const (
	GetPositionsRequestQueryProductAll      GetPositionsRequestQueryProduct = "0"
	GetPositionsRequestQueryProductGenbutsu GetPositionsRequestQueryProduct = "1"
	GetPositionsRequestQueryProductMargin   GetPositionsRequestQueryProduct = "2"
)

type getPositionsRequestOptions struct {
	Product *GetPositionsRequestQueryProduct
}

func WithGetPositionsRequestQuery(q GetPositionsRequestQueryProduct) GetPositionsRequestOption {
	return func(o *getPositionsRequestOptions) {
		o.Product = &q
	}
}

type GetPositionsResponse struct {
	// 約定番号
	ExecutionID string `json:"ExecutionID"`
	// 口座種別
	AccountType AccountType `json:"AccountType"`
	// 銘柄コード
	Symbol Symbol `json:"Symbol"`
	// 銘柄名
	SymbolName string `json:"SymbolName"`
	// 市場コード
	Exchange Exchange `json:"Exchange"`
	// 市場名
	ExchangeName string `json:"ExchangeName"`
	// 約定日（建玉日）
	ExecutionDay Date `json:"ExecutionDay"`
	// 値段
	Price float64 `json:"Price"`
	// 残数量
	LeavesQty float64 `json:"LeavesQty"`
	// 拘束数量（保有数量）
	HoldQty float64 `json:"HoldQty"`
	// 売買区分
	Side Side `json:"Side"`
	// 諸経費
	Expenses float64 `json:"Expenses"`
	// 手数料
	Commission float64 `json:"Commission"`
	// 手数料消費税
	CommissionTax float64 `json:"CommissionTax"`
	// 返済期日
	ExpireDay Date `json:"ExpireDay"`
	// 信用取引区分
	MarginTradeType MarginTradeType `json:"MarginTradeType"`
	// 現値
	CurrentPrice float64 `json:"CurrentPrice"`
	// 評価金額
	Valuation float64 `json:"Valuation"`
	// 評価損益額
	ProfitLoss float64 `json:"ProfitLoss"`
	// 評価損益率
	ProfitLossRate float64 `json:"ProfitLossRate"`
}
