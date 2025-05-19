package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config структура для хранения конфигурации
type Config struct {
	APIURL           string `mapstructure:"api_url"`
	DefaultSize      string `mapstructure:"default_size"`
	TelegramBotToken string `mapstructure:"telegram_bot_token"`
}

// LoadConfig загружает конфигурацию из файла
func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл конфигурации: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("не удалось распарсить конфигурацию: %w", err)
	}

	return &config, nil
}
