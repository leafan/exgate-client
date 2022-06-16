package trade

import (
	"time"
)

// Strategy action.
type Action int

const (
	SELL Action = -1
	HOLD Action = 0
	BUY  Action = 1
)

// Asset values.
type Asset struct {
	Date    []time.Time
	Opening []float64
	Closing []float64
	High    []float64
	Low     []float64
	Volume  []int64
}

// Strategy function. It takes an Asset and returns
// actions for each row.
type StrategyFunction func(Asset) []Action

// Buy and hold strategy
func BuyAndHoldStrategy(asset Asset) []Action {
	actions := make([]Action, len(asset.Date))

	for i := 0; i < len(actions); i++ {
		actions[i] = BUY
	}

	return actions
}
