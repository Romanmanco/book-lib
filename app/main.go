// Package main Author: Роман Манько [@speakerkiller]
package main

import (
	"book-lib/config"
	_ "book-lib/docs"
	"book-lib/internal/api"
	"book-lib/internal/service"
	"book-lib/internal/storage"
	"book-lib/logger"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
)

// Package main Book Library API.
//
//	@title			Book Library API
//	@version		1.0
//	@description	REST API для управления библиотекой книг.
//	@termsOfService	http://swagger.io/terms/
//	@host			localhost:8080
//	@BasePath		/
//	@schemes		http
func main() {
	// Переменные окружения из .env
	config.LoadEnv()

	// Инициализация логгера
	logger.InitLogger()
	logger.Info("Запуск сервера...")

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

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Запуск сервера
	port := config.GetEnv("PORT", "8080")
	log.Printf("Сервер запущен на :%s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
