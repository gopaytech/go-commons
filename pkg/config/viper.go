package config

import (
	"github.com/gopaytech/go-commons/pkg/util"
	"github.com/spf13/viper"
	"strings"
)

func NewConfig(configName string, configPath string, prefix string) *viper.Viper {
	fang := viper.New()

	if !util.IsStringEmpty(prefix) {
		fang.SetEnvPrefix(prefix)
	}
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()

	fang.SetConfigName(configName)
	fang.AddConfigPath(".")
	fang.AddConfigPath(configPath)
	_ = fang.ReadInConfig()

	return fang
}
