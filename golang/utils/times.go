package utils

type Period int64

func (p Period) Prev(ts int64, multiple int64) int64 {
	pv := int64(p)
	return (ts/pv)*pv - pv*multiple
}

func (p Period) Current(ts int64) int64 {
	pv := int64(p)
	return (ts / pv) * pv
}

func (p Period) Next(ts int64, multiple int64) int64 {
	pv := int64(p)
	return (ts/pv)*pv + pv*multiple
}
