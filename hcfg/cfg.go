package hcfg

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"sync"
	"time"
)

func GetString(key string, def ...string) string {
	ok := regUse(key)
	if ok {
		return viper.GetString(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

func MustGetString(key string) (string, error) {
	val := GetString(key)
	if len(val) == 0 {
		return val, errors.New(fmt.Sprintf("miss config: %s ", key))
	}
	return val, nil
}

func GetBool(key string, def ...bool) bool {
	ok := regUse(key)
	if ok {
		return viper.GetBool(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return false
}

func GetInt(key string, def ...int) int {
	ok := regUse(key)
	if ok {
		return viper.GetInt(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func GetInt64(key string, def ...int64) int64 {
	ok := regUse(key)
	if ok {
		return viper.GetInt64(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func GetUint64(key string, def ...uint64) uint64 {
	ok := regUse(key)
	if ok {
		return viper.GetUint64(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func GetDuration(key string, def ...time.Duration) time.Duration {
	ok := regUse(key)
	if ok {
		return viper.GetDuration(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

var gRegisteredCfgMap = make(map[string]any)
var gRegisteredCfgMu sync.Mutex

// Dump Used to output all used configurations
func Dump(call func(key string, val any)) {
	for k, v := range gRegisteredCfgMap {
		call(k, v)
	}
}

func regUse(key string) bool {
	val := viper.Get(key)
	gRegisteredCfgMu.Lock()
	defer gRegisteredCfgMu.Unlock()
	gRegisteredCfgMap[key] = val
	return val != nil
}
