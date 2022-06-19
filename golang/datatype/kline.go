package datatype

type KlineRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	Limit     int    `json:"limit"`
}

// 1499040000000,      // Open time
// "0.01634790",       // Open
// "0.80000000",       // High
// "0.01575800",       // Low
// "0.01577100",       // Close
// "148976.11427815",  // Volume
// 1499644799999,      // Close time

// [
//   [
//     1499040000000,      // 开盘时间
//     "0.01634790",       // 开盘价
//     "0.80000000",       // 最高价
//     "0.01575800",       // 最低价
//     "0.01577100",       // 收盘价(当前K线未结束的即为最新价)
//     "148976.11427815",  // 成交量
//     1499644799999,      // 收盘时间
//     "2434.19055334",    // 成交额
//     308,                // 成交笔数
//     "1756.87402397",    // 主动买入成交量
//     "28.46694368",      // 主动买入成交额
//     "17928899.62484339" // 请忽略该参数
//   ]
// ]
type Kline struct {
	OpenTime  float64 `json:"openTime"`
	Open      string  `json:"open"`
	High      string  `json:"high"`
	Low       string  `json:"low"`
	Close     string  `json:"close"`
	Volume    string  `json:"volume"`
	CloseTime float64 `json:"closeTime"`
}

type RecentTradeRequest struct {
	Symbol string `json:"symbol"`
}

type RecentTrade struct {
	Id           int64  `json:"id"`
	Time         int64  `json:"time"`
	Price        string `json:"price"`
	Amount       int64  `json:"amount"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}
