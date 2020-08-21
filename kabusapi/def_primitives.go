package kabusapi

import (
	fmt "fmt"
	"strconv"
	"strings"
	"time"
)

// 銘柄コード
type Symbol string

// 市場コード
type Exchange int32

const (
	// 東証
	Tosho Exchange = 1
	// 名証
	Meisho Exchange = 3
	// 福証
	Fukusho Exchange = 5
	// 札証
	Sassho Exchange = 6
)

// 商品種別
type SecurityType int32

const (
	Stock SecurityType = 1
)

// 売買区分
type Side string

const (
	// 売
	Sell Side = "1"
	// 買
	Buy Side = "2"
)

// 現物信用区分
type CashMargin = int32

const (
	// 現物
	Genbutsu CashMargin = 1 // FIXME: rename
	// 信用新規
	ShinyoShinki CashMargin = 2 // FIXME: rename
	// 信用返済
	ShinyoHensai CashMargin = 3 // FIXME: rename
)

// 信用取引区分
type MarginTradeType = int32

const (
	// 制度信用
	SystemMarginTrade MarginTradeType = 1
	// 一般信用
	GeneralMarginTrade MarginTradeType = 2
	// 一般信用（売短）
	GeneralMarginShortSellingTrade MarginTradeType = 3
)

// 受渡区分
type DelivType = int32

const (
	// 指定なし
	DelivTypeUnspecified DelivType = 0
	// 自動振替
	AutomaticTransfer DelivType = 1
	// お預り金
	Deposit DelivType = 2
)

// 資産区分
type FundType string

const (
	// 指定なし
	FundTypeUnspecified FundType = "  "
	// 保護
	Protection FundType = "02" // FIXME: rename
	// 信用代用
	ShinyoSubstitute FundType = "AA" // FIXME: rename
	// 証拠金代用
	ShokokinSubstitute FundType = "BB" // FIXME: rename
	// 信用取引
	MarginTrade FundType = "11" // FIXME: rename
)

// 口座種別
type AccountType int32

const (
	// 一般
	GeneralAccount AccountType = 2
	// 特定
	SpecifiedAccount AccountType = 4
	// 法人
	CorporateAccount AccountType = 12
)

// 注文数量
type Qty int32

// 決済順序
type ClosePositionOrder *int32

var (
	// 日付（古い順）、損益（高い順）
	AscDateDescPL ClosePositionOrder = Int32P(0)
	// 日付（古い順）、損益（低い順）
	AscDateAscPL ClosePositionOrder = Int32P(1)
	// 日付（新しい順）、損益（高い順）
	DescDateDescPL ClosePositionOrder = Int32P(2)
	// 日付（新しい順）、損益（低い順）
	DescDateAscPL ClosePositionOrder = Int32P(3)
	// 損益（高い順）、日付（古い順）
	DescPLAscDate ClosePositionOrder = Int32P(4)
	// 損益（高い順）、日付（新しい順）
	DescPLDescDate ClosePositionOrder = Int32P(5)
	// 損益（低い順）、日付（古い順）
	AscPLAscDate ClosePositionOrder = Int32P(6)
	// 損益（低い順）、日付（新しい順）
	AscPLDescDate ClosePositionOrder = Int32P(7)
)

// 建玉ID
type HoldID string

type ClosePosition struct {
	// 返済建玉ID
	HoldID HoldID `json:"HoldID"`
	// 返済建玉数量
	Qty Qty `json:"Qty"`
}

// 注文価格
type OrderPrice uint32

//genconstructor
type Date struct {
	Year  int32      `required:""`
	Month time.Month `required:""`
	Day   int32      `required:""`
}

func (d Date) MarshalJSON() ([]byte, error) {
	s := strings.TrimLeft(fmt.Sprintf(
		"%d%02d%02d",
		d.Year,
		d.Month,
		d.Day,
	), "0")
	if len(s) == 0 {
		s = "0"
	}
	return []byte(s), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	num, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	d.Year = int32(num / 10000)
	d.Month = time.Month(num / 100 % 100)
	d.Day = int32(num % 100)
	return nil
}

func (d Date) ToTime(loc *time.Location) time.Time {
	return time.Date(int(d.Year), d.Month, int(d.Day), 0, 0, 0, 0, loc)
}

//genconstructor
type DateTime struct {
	time.Time `required:""`
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return []byte(d.Format(`"2006-01-02T15:04:05.999999"`)), nil
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02T15:04:05.999999"`, string(b))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// 執行条件
type FrontOrderType int32

const (
	// 成行
	Nariyuki FrontOrderType = 10
	// 寄成（前場）
	YorinariZenba FrontOrderType = 13
	// 寄成（後場）
	YorinariGoba FrontOrderType = 14
	// 引成（前場）
	HikenariZenba FrontOrderType = 15
	// 引成（後場）
	HikenariGoba FrontOrderType = 16
	// IOC成行
	IOCNariyuki FrontOrderType = 17
	// 指値
	Sashine FrontOrderType = 20
	// 寄指（前場）
	YorisashiZenba FrontOrderType = 21
	// 寄指（後場）
	YorisashiGoba FrontOrderType = 22
	// 引指（前場）
	HikesashiZenba FrontOrderType = 23
	// 引指（後場）
	HikesashiGoba FrontOrderType = 24
	// 不成（前場）
	FunariZenba FrontOrderType = 25
	// 不成（後場）
	FunariGoba FrontOrderType = 26
	// IOC指値
	IOCSashine FrontOrderType = 27
)

// 注文番号
type OrderID string

// 現値前値比較
type CurrentPriceChangeStatus string

const (
	// 事象なし
	CurrentPriceChangeStatusNoEvent CurrentPriceChangeStatus = "0000"
	// 変わらず
	Unchanged CurrentPriceChangeStatus = "0056"
	// UP
	UP CurrentPriceChangeStatus = "0057"
	// DOWN
	Down CurrentPriceChangeStatus = "0058"
	// 中断板寄り後の初値
	InitialPriceAfterApprochingTheSuspesionBoard CurrentPriceChangeStatus = "0059" // FIXME: rename
	// ザラバ引け
	ContinuousSessionClosing CurrentPriceChangeStatus = "0060" // FIXME: rename
	// 板寄り引け
	ItayoriClosing CurrentPriceChangeStatus = "0061" // FIXME: rename
	// 中断引け
	InterruptionClosing CurrentPriceChangeStatus = "0062" // FIXME: rename
	// ダウン引け
	DownCLosing CurrentPriceChangeStatus = "0063" // FIXME: rename
	// 逆転終値
	ReversalPrice CurrentPriceChangeStatus = "0064" // FIXME: rename
	// 特別気配引け
	SpecialQuoteClosing CurrentPriceChangeStatus = "0066" // FIXME: rename
	// 一時留保引け
	IchijiRyuhoClosing CurrentPriceChangeStatus = "0067" // FIXME: rename
	// 売買停止引け
	StopTradingClosing CurrentPriceChangeStatus = "0068" // FIXME: rename
	// サーキットブレーカ引け
	CircuitBreakerClosing CurrentPriceChangeStatus = "0069"
	// ダイナミックサーキットブレーカ引け
	DynamicCircuitBreakerClosing CurrentPriceChangeStatus = "0431"
)

// 現値ステータス
type CurrentPriceStatus int32

const (
	// 現値
	CurrentPrice CurrentPriceStatus = 1
	// 不連続歩み
	DiscontinuousStep CurrentPriceStatus = 2 // FIXME: rename
	// 板寄せ
	Itayose CurrentPriceStatus = 3 // FIXME: rename
	// システム障害
	SystemFailure CurrentPriceStatus = 4
	// 中断
	Interruption CurrentPriceStatus = 5
	// 売買停止
	StopTrading CurrentPriceStatus = 6
	// 売買停止・システム停止解除
	StopTradingAndLiftSystemStop CurrentPriceStatus = 7
	// 終値
	ClosingPrice CurrentPriceStatus = 8
	// システム停止
	SystemStop CurrentPriceStatus = 9
	// 概算値
	ApproximatePrice CurrentPriceStatus = 10
	// 参考値
	ReferencePrice CurrentPriceStatus = 11
	// サーキットブレイク実施中
	CircuitBreakeIsInProgress CurrentPriceStatus = 12
	// システム障害解除
	LiftSystemFailure CurrentPriceStatus = 13
	// サーキットブレイク解除
	LiftCircuitBreake CurrentPriceStatus = 14
	// 中断解除
	LiftInterruption CurrentPriceStatus = 15
	// 一時留保中
	IchijiRyuhoIsInProgress CurrentPriceStatus = 16 // FIXME: rename
	// 一時留保解除
	ReleseIchijiRyoho CurrentPriceStatus = 17 // FIXME: rename
	// ファイル障害
	FileFailure CurrentPriceStatus = 18
	// ファイル障害解除
	LiftFileFailure CurrentPriceStatus = 19
	// Spread/Strategy
	SpreadStorategy CurrentPriceStatus = 20
	// ダイナミックサーキットブレイク発動
	TriggerDynamicCircuitBreake CurrentPriceStatus = 21
	// ダイナミックサーキットブレイク解除
	LiftDynamicCircuitBreake CurrentPriceStatus = 22
	// 板寄せ約定
	ItayoseYakujo CurrentPriceStatus = 23 // FIXME: rename
)

// 最良気配フラグ
type QuoteSign string

const (
	// 事象なし
	QuoteSignNoEvent QuoteSign = "0000"
	// 一般気配
	GenralQuote QuoteSign = "0101"
	// 特別気配
	SpecialQuote QuoteSign = "0102"
	// 注意気配
	AttentionQuote QuoteSign = "0103"
	// 寄前気配
	BeforeOpeningQuote QuoteSign = "0107"
	// 停止前特別気配
	BeforeStoppingQuote QuoteSign = "0108"
	// 引け後気配
	AfterClosingQuote QuoteSign = "0109"
	// 寄前気配約定成立ポイントなし
	BeforeOpeningWithoutPoint QuoteSign = "0116" // FIXME: rename
	// 寄前気配約定成立ポイントあり
	BeforeOpeningWithPoint QuoteSign = "0117" // FIXME: rename
	// 連続約定気配
	ContinuousExecutionQuote QuoteSign = "0118"
	// 停止前の連続約定気配
	ContinuousExecutionQuoteBeforeStopping QuoteSign = "0119" // FIXME: rename
	// 買い上がり売り下がり中
	KaiagariUrisagari QuoteSign = "0120" // FIXME: rename
)

// 業種コード名
type BisCategory string

const (
	// 水産・農林業
	FisheriesAndForestries BisCategory = "0050"
	// 鉱業
	Mining BisCategory = "1050"
	// 建設業
	ConstructionIndustry BisCategory = "2050"
	// 食料品
	Grocery BisCategory = "3050"
	// 繊維製品
	TextileProducts BisCategory = "3100"
	// パルプ・紙
	PulpAndPaper BisCategory = "3150"
	// 化学
	Chemistry BisCategory = "3200"
	// 医薬品
	Pharmaceuticals BisCategory = "3250"
	// 石油・石炭製品
	OilAndCoalProducts BisCategory = "3300"
	// ゴム製品
	RubberProducts BisCategory = "3350"
	// ガラス・土石製品
	GlassAndStoneProducts BisCategory = "3400"
	// 鉄鋼
	Steel BisCategory = "3450"
	// 非鉄金属
	NonFerrousMetal BisCategory = "3500"
	// 金属製品
	MetalProducts BisCategory = "3550"
	// 機械
	Machine BisCategory = "3600"
	// 電気機器
	ElectricalEquipment BisCategory = "3650"
	// 輸送用機器
	TransportEquipment BisCategory = "3700"
	// 精密機器
	PrecisionEquipment BisCategory = "3750"
	// その他製品
	OtherProducts BisCategory = "3800"
	// 電気・ガス業
	ElectricityAndGasIndustry BisCategory = "4050"
	// 陸運業
	LandTransportation BisCategory = "5050"
	// 海運業
	Shipping BisCategory = "5100"
	// 空運業
	AritTransportation BisCategory = "5150"
	// 倉庫・運輸関連業
	WarehouseAndTransportationRelatedBusiness BisCategory = "5200"
	// 情報・通信業
	InformationAndCommunicationIndustry BisCategory = "5250"
	// 卸売業
	Wholesale BisCategory = "6050"
	// 小売業
	Retail BisCategory = "6100"
	// 銀行業
	Banking BisCategory = "7050"
	// 証券、商品先物取引業
	SecuritiesAndCommodityFuturesTradingBusiness BisCategory = "7100"
	// 保険業
	InsuranceIndustry BisCategory = "7150"
	// その他金融業
	OtherFinancialServices BisCategory = "7200"
	// 不動産業
	RealEstateIndustry BisCategory = "8050"
	// サービス業
	ServiceIndustry BisCategory = "9050"
	// その他
	Other BisCategory = "9999"
)

// 呼値コード
type PriceRangeCode string

const (
	PriceRangeCode10000 PriceRangeCode = "10000"
	PriceRangeCode10003 PriceRangeCode = "10003"
)

// 注文状態
type OrderState int32

const (
	// 待機（発注待機）
	WaitingForOrder OrderState = 1
	// 処理中（発注送信中）
	SendingOrder OrderState = 2
	// 処理済（発注済・訂正済）
	Ordered OrderState = 3
	// 訂正取消送信中
	SendingCancellation OrderState = 4
	// 終了（発注エラー・取消済・全約定・失効・期限切れ）
	EndOrder OrderState = 5
)

// 執行条件
type OrderType int32

const (
	// ザラバ
	Zaraba OrderType = 1
	// 寄り
	Yori OrderType = 2
	// 引け
	Hike OrderType = 3
	// 不成
	Funari OrderType = 4
)

// 注文明細種別
type RecType int32

const (
	// 受付
	RecTypeReceived RecType = 1
	// 繰越
	RecTypeCarryover RecType = 2
	// 期限切れ
	RecTypeExpired RecType = 3
	// 発注
	RecTypeOrder RecType = 4
	// 訂正
	RecTypeCorrection RecType = 5
	// 取消
	RecTypeCancel RecType = 6
	// 失効
	RecTypeRevocation RecType = 7
	// 約定
	RecTypeExecution RecType = 8
)

// 注文状態
type OrderDetailState int32

const (
	// 待機（発注待機）
	OrderDetailStateWaiting OrderDetailState = 1
	// 処理中（発注送信中・訂正送信中・取消送信中）
	OrderDetailStateInProcess OrderDetailState = 2
	// 処理済（発注済・訂正済・取消済・全約定・期限切れ）
	OrderDetailStateProcessed OrderDetailState = 3
	// エラー
	OrderDetailStateError OrderDetailState = 4
	// 削除済み
	OrderDetailStateDeleted OrderDetailState = 5
)
