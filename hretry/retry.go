package hretry

import (
	"github.com/avast/retry-go"
	"github.com/hootuu/hyle/hcfg"
	"github.com/hootuu/hyle/hlog"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"time"
)

func Universal(call func() error) {
	err := retry.Do(func() error {
		return call()
	},
		retry.Attempts(cast.ToUint(hcfg.GetInt("retry.universal.attempts", 3))),
		retry.Delay(hcfg.GetDuration("retry.universal.delay", 600*time.Millisecond)),
	)
	if err != nil {
		hlog.Err("hyle.retry.Universal", zap.Error(err))
	}
}

func Must(call func() error) error {
	err := retry.Do(func() error {
		return call()
	},
		retry.Attempts(cast.ToUint(hcfg.GetInt("retry.must.attempts", 3))),
		retry.Delay(hcfg.GetDuration("retry.must.delay", 600*time.Millisecond)),
	)
	if err != nil {
		hlog.Err("hyle.retry.Universal", zap.Error(err))
		return err
	}
	return nil
}
