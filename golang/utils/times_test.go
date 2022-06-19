package utils

import (
	"testing"
	"time"
)

func Test_Ts(t *testing.T) {
	ts := time.Now().Unix()
	t.Log(ts)
	m := Period(300)
	t.Log(ts, m.Prev(ts, 1), m.Current(ts), m.Next(ts, 1))
	m = Period(60 * 60)
	t.Log(ts, m.Prev(ts, 1), m.Current(ts), m.Next(ts, 1))
}
