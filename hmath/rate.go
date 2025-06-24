package hmath

import (
	"fmt"
	"github.com/hootuu/hyle/hlog"
	"go.uber.org/zap"
	"math"
)

type Rate uint64

type RateBase int

const (
	Rate10         RateBase = 1
	Rate100        RateBase = 2
	Rate1000       RateBase = 3
	Rate10000      RateBase = 4
	Rate100000     RateBase = 5
	Rate1000000    RateBase = 6
	Rate10000000   RateBase = 7
	Rate100000000  RateBase = 8
	Rate1000000000 RateBase = 9
)

func (r RateBase) Validate() bool {
	switch r {
	case Rate10, Rate100, Rate1000,
		Rate10000, Rate100000, Rate1000000,
		Rate10000000, Rate100000000, Rate1000000000:
		return true
	}
	return false
}

func NewRate(b RateBase, value uint64) Rate {
	if ok := b.Validate(); !ok {
		hlog.Err("invalid RateBase", zap.Int("base", int(b)))
		b = Rate10000
	}
	return Rate(value*10 + uint64(b))
}

func (r Rate) Base() RateBase {
	b := int(r % 10)
	return RateBase(b)
}

func (r Rate) Value() uint64 {
	v := uint64(r / 10)
	return v
}

func (r Rate) Rate() float64 {
	v := float64(r.Value())
	b := math.Pow10(int(r.Base()))
	return v / b
}

func (r Rate) Str() string {
	v := r.Rate()
	return fmt.Sprintf("%.2f", v)
}

func (r Rate) String(n int) string {
	format := fmt.Sprintf("%%.%df", n)
	return fmt.Sprintf(format, r.Value())
}
