package trade

func ApplyActions(prices []float64, actions []Action) []float64 {
	gains := make([]float64, len(actions))

	initialBalance := 1.0
	balance := initialBalance
	shares := 0.0

	for i := 0; i < len(actions); i++ {
		switch actions[i] {
		case BUY:
			if balance > 0 {
				shares = balance / prices[i]
				balance = 0
			}

		case SELL:
			if shares > 0 {
				balance = shares * prices[i]
				shares = 0
			}
		}

		gains[i] = ((shares * prices[i]) + balance - initialBalance) / initialBalance
	}

	return gains
}
