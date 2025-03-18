// Package main Author: Роман Манько [@speakerkiller]
package main

import (
	"book-lib/internal/api"
	"book-lib/internal/service"
	"book-lib/internal/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// хранилище и сервис для работы с книгами
	store := storage.NewBookStorage()
	bookService := service.NewBookService(store)

	// маршруты
	api.SetupRoutes(e, bookService)

	// запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
