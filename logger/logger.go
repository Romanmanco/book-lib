// Package logger Author: Роман Манько [@speakerkiller]
package logger

import (
	"log"
	"os"
)

// LogLevel уровень логирования (DEBUG, INFO)
var LogLevel = "INFO"

// InitLogger инициализация уровня логирования
func InitLogger() {
	if level, exists := os.LookupEnv("LOG_LEVEL"); exists {
		LogLevel = level
	}
}

// Debug логирование уровня DEBUG
func Debug(v ...interface{}) {
	if LogLevel == "DEBUG" {
		log.Println("[DEBUG]", v)
	}
}

// Info логирование уровня INFO
func Info(v ...interface{}) {
	log.Println("[INFO]", v)
}

// Error логирование ошибок
func Error(v ...interface{}) {
	log.Println("[ERROR]", v)
}
