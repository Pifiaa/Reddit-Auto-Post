package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitializeConfig() (bool, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()

	if err != nil {
		return false, fmt.Errorf("Error leyendo el archivo de configuración, %s", err)
	}

	return true, nil
}

/**
 *
 */
func GetEnv(key string) string {
	return viper.GetString(key)
}
