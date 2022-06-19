package utils

import (
	"fmt"
	"strconv"
)

func Float64ToString(f float64) string {
	return fmt.Sprintf("%.8f", f)
}

func StringTofloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func Float64ToInt64(f float64) int64 {
	return int64(f)
}
