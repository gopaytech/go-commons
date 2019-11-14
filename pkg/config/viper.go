package config

import (
	"github.com/gopaytech/go-commons/pkg/encoding"
	strings2 "github.com/gopaytech/go-commons/pkg/strings"
	"github.com/spf13/viper"
	"strings"
)

func NewConfig(configName string, configPath string, prefix string) *viper.Viper {
	fang := viper.New()

	if !strings2.IsStringEmpty(prefix) {
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

func GetStringDefault(viper *viper.Viper, key string, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func GetArrayString(viper *viper.Viper, key string) []string {
	return GetStringSplit(viper, key, ",")
}

func GetStringSplit(viper *viper.Viper, key string, separator string) []string {
	return strings.Split(viper.GetString(key), separator)
}

func GetDecodeBase64(viper *viper.Viper, key string) (plain string, err error) {
	encodedValue := viper.GetString(key)
	return encoding.Base64Decode(encodedValue)
}

func GetDecodeBase32(viper *viper.Viper, key string) (plain string, err error) {
	encodedValue := viper.GetString(key)
	return encoding.Base32Decode(encodedValue)
}
