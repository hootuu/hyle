package hcfg

import (
	"github.com/spf13/viper"
	"strings"
)

func SetEvnPrefix(in string) {
	viper.SetEnvPrefix(in)
}

func AddConfigPath(path []string) {
	for _, p := range path {
		viper.AddConfigPath(p)
	}
}

func SetConfigName(name string) {
	viper.SetConfigName(name)
}

func SetConfigType(t string) {
	viper.SetConfigType(t)
}

func ReadInConfig() {
	_ = viper.ReadInConfig()
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
