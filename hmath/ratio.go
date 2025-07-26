package hmath

type Ratio struct {
	Base  uint64 `json:"base"`
	Ratio uint64 `json:"ratio"`
}

func NewRatio(base uint64, ratio uint64) *Ratio {
	if base == 0 {
		base = 100
	}
	if ratio == 0 {
		ratio = 100
	}
	return &Ratio{
		Base:  base,
		Ratio: ratio,
	}
}

func (r *Ratio) Mul(val uint64) float64 {
	return float64(val) * float64(r.Ratio) / float64(r.Base)
}

func (r *Ratio) Div(val uint64) float64 {
	return float64(val) * float64(r.Base) / float64(r.Ratio)
}
