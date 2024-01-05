package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App    App
		Server Server
		Db     Db
		Reddit Reddit
	}

	App struct {
		Name string `mapstructure:"Name"`
	}

	Server struct {
		Host string `mapstructure:"Host"`
		Port string `mapstructure:"Port"`
	}

	Db struct {
		Host     string `mapstructure:"Host"`
		Port     string `mapstructure:"Port"`
		User     string `mapstructure:"User"`
		Password string `mapstructure:"Password"`
		DBName   string `mapstructure:"DBName"`
		SSLMode  string `mapstructure:"SSLMode"`
		TimeZone string `mapstructure:"TimeZone"`
	}

	Reddit struct {
		Url   string `mapstructure:"Url"`
		Oauth string `mapstructure:"Oauth"`
	}
)

var (
	initialized bool
	appConfig   *Config
	once        sync.Once
	mu          sync.Mutex
)

func loadConfig() error {
	rootDir, err := getRootDirectory()
	if err != nil {
		return err
	}

	configureViper(rootDir)

	if err := readConfig(); err != nil {
		return err
	}

	appConfig = &Config{}

	err = viper.Unmarshal(appConfig)

	if err != nil {
		return err
	}

	return nil
}

func getRootDirectory() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", fmt.Errorf("Error al obtener el directorio actual: %w", err)
	}

	return filepath.Join(dir, "..", ".."), nil
}

func configureViper(rootDir string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(rootDir)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
}

func readConfig() error {
	if err := viper.ReadInConfig(); err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return fmt.Errorf("Error al leer el archivo de configuración: %w", err)
	}
	return nil
}

func InitConfig() error {
	mu.Lock()
	defer mu.Unlock()

	if !initialized {
		once.Do(func() {
			err := loadConfig()

			if err != nil {
				fmt.Println("Error loading config:", err)
			}

			initialized = true
		})
	}

	return nil
}

func GetConfig() *Config {
	return appConfig
}
