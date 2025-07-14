package hcoin

import (
	"github.com/hootuu/hyle/hlog"
	"go.uber.org/zap"
)

type Currency string
type Decimals = uint8

const (
	CNY Currency = "CNY"
	USD Currency = "USD"
	HKD Currency = "HKD"
)

func (c Currency) String() string {
	return string(c)
}

var gCurrencyDict = map[Currency]Decimals{
	CNY: Decimals(2),
	USD: Decimals(2),
	HKD: Decimals(2),
}

func (c Currency) Decimals() Decimals {
	d, ok := gCurrencyDict[c]
	if !ok {
		hlog.Err("hyle.coin: no such currency", zap.String("currency", string(c)))
		return 0
	}
	return d
}

func RegisterCurrency(c Currency, d Decimals) {
	gCurrencyDict[c] = d
}
