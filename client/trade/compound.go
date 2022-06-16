package trade

// MACD and RSI strategy.
func MacdAndRsiStrategy(asset Asset, fastN, slowN, singnal int64) []Action {
	strategy := AllStrategies(MacdStrategy, DefaultRsiStrategy)
	return strategy(asset)
}

func AllStrategies(strategies ...StrategyFunction) StrategyFunction {
	return func(asset Asset) []Action {
		actions := RunStrategies(asset, strategies...)

		for i := 1; i < len(actions); i++ {
			for j := 0; j < len(actions[0]); j++ {
				if actions[0][j] != actions[i][j] {
					actions[0][j] = HOLD
				}
			}
		}

		return actions[0]
	}
}

func RunStrategies(asset Asset, strategies ...StrategyFunction) [][]Action {
	actions := make([][]Action, len(strategies))

	for i := 0; i < len(strategies); i++ {
		actions[i] = strategies[i](asset)
	}

	return actions
}

//todo: add optimize rsi strategy
func DefaultRsiStrategy(asset Asset) []Action {

	return RsiStrategy(asset, 70, 30, 14)
}
