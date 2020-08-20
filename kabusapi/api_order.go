package kabusapi

import "context"

func (c Client) SendOrder(
	ctx context.Context,
	req SendOrderRequest,
) (SendOrderResponse, error) {
	res := SendOrderResponse{}
	err := c.postRequest(ctx, "/sendorder", req, &res)
	return res, err
}

// 注文パスワード
type OrderPassword string

type SendOrderRequest struct {
	// 注文パスワード
	Password OrderPassword `json:"Password"`
	// 銘柄コード
	Symbol Symbol `json:"Symbol"`
	// 事情コード
	Exchange Exchange `json:"Exchange"`
	// 商品種別
	SecurityType SecurityType `json:"SecurityType"`
	// 売買区分
	Side Side `json:"Side"`
	// 現物信用区分
	CashMargin CashMargin `json:"CashMargin"`
	// 信用取引区分
	MarginTradeType MarginTradeType `json:"MarginTradeType"`
	// 受渡区分
	DelivType DelivType `json:"DelivType"`
	// 資産区分
	FundType FundType `json:"FundType"`
	// 口座種別
	AccountType AccountType `json:"AccountType"`
	// 注文数量
	Qty Qty `json:"Qty"`
	// 決済順序
	ClosePositionOrder ClosePositionOrder `json:"ClosePositionOrder"`
	// 返済建玉指定
	ClosePositions []ClosePosition `json:"ClosePositions"`
	// 注文価格
	Price OrderPrice `json:"Price"`
	// 注文有効期限
	ExpireDay Date `json:"ExpireDay"`
	// 執行条件
	FrontOrderType FrontOrderType `json:"FrontOrderType"`
}

type SendOrderResponse struct {
	// 受付注文番号
	OrderID OrderID `json:"OrderId"`
}

func (c Client) CancelOrder(
	ctx context.Context,
	orderID OrderID,
	password OrderPassword,
) (CancelOrderResponse, error) {
	req := NewCancelOrderRequest(orderID, password)
	res := CancelOrderResponse{}
	err := c.putRequest(ctx, "/cancelorder", req, &res)
	return res, err
}

//genconstructor
type CancelOrderRequest struct {
	// 注文番号
	OrderID OrderID `json:"OrderId" required:""`
	// 注文パスワード
	Password OrderPassword `json:"Password" required:""`
}

type CancelOrderResponse struct {
	// 受付注文番号
	OrderID OrderID `json:"OrderId"`
}
