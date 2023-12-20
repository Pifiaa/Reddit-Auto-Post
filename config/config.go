package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// InitializeConfig inicializa la configuración desde el archivo "config.yml".
//
// La función busca el archivo de configuración "config.yml" en el directorio
// raíz del proyecto. Si no se encuentra, devuelve un error indicando la ausencia
// del archivo. Si hay un error al leer el archivo de configuración, también
// devuelve un error.
func InitializeConfig() error {
	dir, _ := os.Getwd()
	rootDir := filepath.Join(dir, "..", "..")

	viper.SetConfigName("config")
	viper.AddConfigPath(rootDir)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("No se encontró el archivo de configuración 'config.yml': %s", err)
		} else {
			return fmt.Errorf("Error al leer el archivo de configuración: %s", err)
		}
	}

	return nil
}

// GetEnv devuelve el valor de una variable de entorno configurada
func GetEnv(key string) string {
	return viper.GetString(key)
}
