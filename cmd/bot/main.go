package main

import (
	"BotQR/internal/bot" // Импортируем пакет bot
	"BotQR/internal/config"
	"log"
)

func main() {
	// Загружаем конфиг
	cfg, err := config.LoadConfig("internal/config")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Запуск бота
	bot.StartBot(cfg)
}
