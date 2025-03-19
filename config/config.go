// Package config Author: Роман Манько [@speakerkiller]
package config

import (
	"book-lib/logger"
	"github.com/joho/godotenv"
	"os"
)

// LoadEnv загружает переменные окружения из .env файла
func LoadEnv() {
	logger.Debug("Загрузка переменных окружения...")
	if err := godotenv.Load("config/.env"); err != nil {
		logger.Error("Не удалось загрузить .env, используются стандартные переменные окружения...")
	}
}

// GetEnv возвращает значение переменной окружения или значение по умолчанию
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
