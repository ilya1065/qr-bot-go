package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// QRClient хранит базовый URL и размер для запросов
type QRClient struct {
	baseURL string
	size    string
}

// NewQRClient создает новый экземпляр QRClient с переданным baseURL и размером
func NewQRClient(baseURL, defaultSize string) *QRClient {
	return &QRClient{
		baseURL: baseURL,
		size:    defaultSize,
	}
}

// GetQRImageURL формирует URL для получения QR-кода
func (c *QRClient) GetQRImageURL(data string) (string, error) {
	if data == "" {
		return "", fmt.Errorf("данные для QR не могут быть пустыми")
	}
	encodedData := url.QueryEscape(data)
	fullURL := fmt.Sprintf("%s?data=%s&size=%s", c.baseURL, encodedData, c.size)

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("ошибка при запросе к API: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API вернул статус %d", resp.StatusCode)
	}
	return fullURL, nil
}

// GenerateQRCode генерирует QR код из данных и возвращает его в виде байтов
func GenerateQRCode(apiURL, size, data string) ([]byte, error) {
	encoded := url.QueryEscape(data)
	fullURL := fmt.Sprintf("%s?data=%s&size=%s", apiURL, encoded, size)

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе к QR API: %w", err)
	}
	defer resp.Body.Close()

	qrData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении тела ответа: %w", err)
	}
	return qrData, nil
}
