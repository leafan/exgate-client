package datatype

type Balances struct {
	Coin   string
	Free   string
	Frozen string
}

type BalancesReq struct {
	Coin string //币种
}

type PendingOrdersRequest struct {
	Symbol  string
	Options map[string]string
}

type PendingOrder struct {
	OrderId string
	Symbol  string
	Type    string
	Price   string
	Amount  string
	Side    string
}

//binance现货交易options
//notice" recived window ?
type BinancePlaceOrderRequestOptions struct {
	RecivedWindow    int64  `json:"recivedWindow"`
	TimeInForce      string `json:"timeInForce"`      // 详见枚举定义：有效方式
	Quantity         int64  `json:"quantity"`         // 数量
	QuoteOrderQty    int64  `json:"quoteOrderQty"`    // 报价数量
	Price            int64  `json:"price"`            // 价格
	NewClientOrderId string `json:"newClientOrderId"` // 客户自定义的唯一订单ID。 如果未发送，则自动生成
	StopPrice        int64  `json:"stopPrice"`        // 仅 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, 和TAKE_PROFIT_LIMIT 需要此参数。
	IcebergQty       int64  `json:"icebergQty"`       // 仅使用 LIMIT, STOP_LOSS_LIMIT, 和 TAKE_PROFIT_LIMIT 创建新的 iceberg 订单时需要此参数
	NewOrderRespType string `json:"newOrderRespType"` // 设置响应JSON。 ACK，RESULT或FULL； "MARKET"和" LIMIT"订单类型默认为"FULL"，所有其他订单默认为"ACK"。
}

//FTX现货交易options
type FtxPlaceOrderRequestOptions struct {
	ReduceOnly bool `json:"reduceOnly"`
	Ioc        bool `json:"ioc"`
	PostOnly   bool `json:"postOnly"`
}

//现货交易
type PlaceOrderRequest struct {
	Symbol        string
	Side          string
	Price         string
	Type          string
	Amount        string
	ClientOrderId string
	Options       map[string]string
}

//期货交易
type FuturePlaceOrderRequest struct {
	Symbol   string
	Side     string
	Price    string
	Type     string
	Quantity string
	Options  map[string]string
}

//现货交易撤单
type CancelOrderRequest struct {
	Symbol        string
	OrderId       string
	ClientOrderId string
	Options       map[string]string
}

type PlaceOrder struct {
	OrderId             string
	NewClientOrderId    string // 系统的订单ID
	Transact_time       string // 交易的时间戳
	Price               string // 订单价格
	CummulativeQuoteQty string // 累计交易的金额
	Type                string // 订单类型， 比如市价单，现价单等
	Side                string // 订单方向，买还是卖
}

type FuturePlaceOrder struct {
	OrderId       string
	ClientOrderId string
}

type FtxNewOrder struct {
	Market     string      `json:"market"`
	Side       string      `json:"side"`
	Price      float64     `json:"price"`
	Type       string      `json:"type"`
	Size       float64     `json:"size"`
	ReduceOnly bool        `json:"reduceOnly"`
	Ioc        bool        `json:"ioc"`
	PostOnly   bool        `json:"postOnly"`
	ClientId   interface{} `json:"clientId"`
}

type BinanceCancelOrderOptions struct {
	RecvWindow int64 `json:"recvWindow"`
	//origClientOrderId string `json:"origClientOrderId"`
	NewClientOrderId string `json:"newClientOrderId"`
}

type FtxCancelOderOptions struct {
}

type CancelOrder struct {
	Symbol string
	Price  string
	Status string
	Side   string
}

type FutureCancelOrder struct {
	Symbol            string
	OrderId           string //系统订单号
	OrigClientOrderId string //用户自定义的订单号
}

type FutureCancelOrderResponse struct {
	Price   string
	CumQty  string
	OrigQty string
	Side    string
	Status  string
}

type FtxPendingOderOptions struct {
	Type string `json:"type"`
}

type BinancePendingOrderOptions struct {
	Timestamp  float64 `json:"timestamp"`
	RecvWindow float64 `json:"recvWindow"`
}

type BinancePlaceOrderResponse struct {
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

// full respone
type FutureBinancePlaceOderResponse struct {
	Symbol            string `json:"symbol"`
	OrderID           int64  `json:"orderId"`
	ClientOrderID     string `json:"clientOrderId"`
	Price             string `json:"price"`
	OrigQuantity      string `json:"origQty"`
	ExecutedQuantity  string `json:"executedQty"`
	CumQuote          string `json:"cumQuote"`
	ReduceOnly        bool   `json:"reduceOnly"`
	Status            string `json:"status"`
	StopPrice         string `json:"stopPrice"`
	TimeInForce       string `json:"timeInForce"`
	Type              string `json:"type"`
	Side              string `json:"side"`
	UpdateTime        int64  `json:"updateTime"`
	WorkingType       string `json:"workingType"`
	ActivatePrice     string `json:"activatePrice"`
	PriceRate         string `json:"priceRate"`
	AvgPrice          string `json:"avgPrice"`
	PositionSide      string `json:"positionSide"`
	ClosePosition     bool   `json:"closePosition"`
	PriceProtect      bool   `json:"priceProtect"`
	RateLimitOrder10s string `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string `json:"rateLimitOrder1m,omitempty"`
}
