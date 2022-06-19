package utils

import (
	tart "github.com/iamjinlei/go-tart"
)

func GenerateMacd(data []float64, fastPeriod int64, slowPeriod int64, signal int64) ([]float64, []float64, []float64) {
	return tart.MacdArr(data, fastPeriod, slowPeriod, signal)
}

func GeneraterRsi(data []float64, period int64) []float64 {
	return tart.RsiArr(data, period)
}
