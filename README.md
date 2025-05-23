# QR-бот на Go

## Описание
  
  Телеграм-бот, который принимает текст и возвращает QR-код, используя внешний API для генерации QR и Telegram API для общения с пользователями. Проект собран на языке Go с возможностью запуска через Docker.

## Основные возможности

- Принимает текстовые сообщения от пользователей
- Генерирует QR-код для каждого сообщения
- Отправляет QR-код обратно в чат Telegram

## Быстрый старт
  
  1. **Клонировать репозиторий:**
  ```bash
  git clone https://github.com/ilya1065/qr-bot-go.git
  cd qr-bot-go
  ```
  
  2. **Собрать и запустить Docker-контейнер:**
  ```bash
  docker build -t botqr .
  docker run -v "ПУТЬ_ДО_ПАПКИ/internal/config:/app/internal/config" botqr
  ```

3. **Файл конфигурации** (пример: `internal/config/config.yaml`):
                                     ```yaml
                                     api_url: "http://api.qrserver.com/v1/create-qr-code/"
                                     default_size: "400x400"
                                     telegram_bot_token: "ВАШ_ТОКЕН_ТЕЛЕГРАМ"
                                     ```
  
  ## CI/CD

                                   - Автоматическая сборка и тесты реализованы через [GitHub Actions](.github/workflows/go-docker.yml).
                                   - Каждый push запускает сборку, тесты и сборку Docker-образа.
  
  ## Структура проекта

                                   - `cmd/bot/` — точка входа приложения
                                   - `internal/api/` — обработчики работы с внешними API
                                   - `internal/bot/` — логика работы с Telegram API
                                   - `internal/config/` — конфигурация и файлы настроек

## Автор
  
  [ilya1065](https://github.com/ilya1065)
