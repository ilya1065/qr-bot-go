# Первый этап: билдим бинарник Go (на большом образе)
FROM golang:1.23 as builder

WORKDIR /app

# Копируем весь проект внутрь контейнера
COPY . .

# Скачиваем зависимости
RUN go mod tidy

# Собираем статически слинкованный бинарник (не зависит от GLIBC!)
RUN CGO_ENABLED=0 GOOS=linux go build -o bot ./cmd/bot

# Второй этап: минимальный образ только с нашим бинарником и файлами конфига
FROM debian:bullseye-slim

WORKDIR /app

# УСТАНАВЛИВАЕМ CA-CERTIFICATES для HTTPS
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Копируем только бинарник из первого этапа
COPY --from=builder /app/bot .
# Копируем конфиги и внутренние файлы (если нужны)
COPY --from=builder /app/internal /app/internal
COPY --from=builder /app/internal/config /app/internal/config

# Запускать наш бот при старте контейнера
CMD ["./bot"]
