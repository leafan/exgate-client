package trade

import "exgate-client/utils"

// MACD strategy.
func MacdStrategy(asset Asset) []Action {
	actions := make([]Action, len(asset.Date))

	macd, signal, _ := utils.GenerateMacd(asset.Closing, 12, 26, 9)

	for i := 0; i < len(actions); i++ {
		if macd[i] > signal[i] {
			actions[i] = BUY
		} else if macd[i] < signal[i] {
			actions[i] = SELL
		} else {
			actions[i] = HOLD
		}
	}

	return actions
}
