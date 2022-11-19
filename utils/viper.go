package utils

import (
	"os"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return os.Getenv(key)
	}

	return viper.GetString(key)
}
