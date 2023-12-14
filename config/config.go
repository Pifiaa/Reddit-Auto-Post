package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// InitializeConfig inicializa la configuración desde el archivo "config.yml"
func InitializeConfig() error {
	dir, _ := os.Getwd()
	rootDir := filepath.Join(dir, "..", "..")

	viper.SetConfigName("config")
	viper.AddConfigPath(rootDir)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()

	if err != nil {
		return fmt.Errorf("Error leyendo el archivo de configuración: %s", err)
	}

	return nil
}

// GetEnv devuelve el valor de una variable de entorno configurada
func GetEnv(key string) string {
	return viper.GetString(key)
}
