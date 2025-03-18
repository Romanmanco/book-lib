// Package main Author: Роман Манько [@speakerkiller]
package main

import (
	"book-lib/config"
	"book-lib/internal/api"
	"book-lib/internal/service"
	"book-lib/internal/storage"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	// Переменные окружения из .env
	config.LoadEnv()

	// Инициализация бд
	db, err := storage.ConnectDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// Миграции
	storage.RunMigrations(db)

	e := echo.New()

	// Сервис работы с книгами
	bookService := service.NewBookService(storage.NewBookStorage(db))

	// Маршруты
	api.SetupRoutes(e, bookService)

	// Запуск сервера
	port := config.GetEnv("PORT", "8080")
	log.Printf("Сервер запущен на :%s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
