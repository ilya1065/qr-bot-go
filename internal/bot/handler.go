package bot

import (
	"BotQR/internal/api"
	"BotQR/internal/config"
	telegramapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func StartBot(cfg *config.Config) {
	bot, err := telegramapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Fatalf("ошибка инициализации бота: %v", err)
	}

	u := telegramapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u) // получаем обновления
	var processedCount int
	for update := range updates {
		if update.Message == nil {
			continue
		}
		text := update.Message.Text
		start := time.Now()
		qr, err := api.GenerateQRCode(cfg.APIURL, cfg.DefaultSize, text)
		if err != nil {
			log.Printf("Ошибка генерации QR: %v", err)
			continue
		}

		msg := telegramapi.NewPhoto(update.Message.Chat.ID, telegramapi.FileBytes{
			Name:  "qr.png",
			Bytes: qr,
		})
		elapsed := time.Since(start)
		bot.Send(msg)
		log.Printf("QR-код успешно отправлен. Текст: %s, Пользователь: %d", text, update.Message.From.ID)
		processedCount++
		log.Printf("Всего обработано сообщений: %d", processedCount)
		log.Printf("Время обработки запроса: %s", elapsed)
		log.Printf("")

	}
}
