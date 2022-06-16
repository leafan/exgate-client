package trade

import "exgate-client/utils"

// RSI strategy. Sells above sell at, buys below buy at.
func RsiStrategy(asset Asset, sellAt, buyAt float64, period int64) []Action {
	actions := make([]Action, len(asset.Date))

	rsi := utils.GeneraterRsi(asset.Closing, period)

	for i := 0; i < len(actions); i++ {
		if rsi[i] <= buyAt {
			actions[i] = BUY
		} else if rsi[i] >= sellAt {
			actions[i] = SELL
		} else {
			actions[i] = HOLD
		}
	}

	return actions
}
