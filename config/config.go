package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitializeConfig() (bool, error) {
	dir, _ := os.Getwd()
	rootDir := filepath.Join(dir, "..", "..")

	viper.SetConfigName("config")
	viper.AddConfigPath(rootDir)
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
