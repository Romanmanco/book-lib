// Package config Author: Роман Манько [@speakerkiller]
package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

// LoadEnv загружает переменные окружения из .env файла
func LoadEnv() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Error("Не удалось загрузить .env, используются стандартные переменные окружения...")
	}
}

// GetEnv возвращает значение переменной окружения или значение по умолчанию
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
