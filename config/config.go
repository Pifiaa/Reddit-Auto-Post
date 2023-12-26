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
		Name string
	}

	Server struct {
		Host string
		Port string
	}

	Db struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}

	Reddit struct {
		Url   string
		Oauth string
	}
)

var (
	initialized bool
	AppConfig   *Config
	once        sync.Once
	mu          sync.Mutex
)

func LoadConfig() (AppConfig *Config, err error) {
	dir, _ := os.Getwd()
	rootDir := filepath.Join(dir, "..", "..")

	viper.SetConfigName("config")
	viper.AddConfigPath(rootDir)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()

	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return nil, fmt.Errorf("Error al leer el archivo de configuración: %w", err)
	}

	AppConfig = &Config{
		App: App{
			Name: viper.GetString("project.name"),
		},

		Server: Server{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},

		Db: Db{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.name"),
			SSLMode:  viper.GetString("database.sslmode"),
			TimeZone: viper.GetString("database.timezone"),
		},

		Reddit: Reddit{
			Url:   viper.GetString("reddit.url"),
			Oauth: viper.GetString("reddit.oauth"),
		},
	}

	return AppConfig, nil
}

func GetConfig() (*Config, error) {
	mu.Lock()
	defer mu.Unlock()

	if !initialized {
		once.Do(func() {
			config, err := LoadConfig()
			if err != nil {
				fmt.Println("Error loading config:", err)
			}
			AppConfig = config
			initialized = true
		})
	}

	return AppConfig, nil
}
