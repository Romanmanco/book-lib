// Package api Author: Роман Манько [@speakerkiller]
package api

import (
	"book-lib/internal/handlers"
	"book-lib/internal/service"
	"github.com/labstack/echo/v4"
)

// SetupRoutes настройка маршрутов
func SetupRoutes(e *echo.Echo, bookService service.BookService) {
	bookHandler := handlers.NewBookHandler(bookService)

	// маршруты для работы с книгами
	e.POST("/books", bookHandler.CreateBook)
	e.GET("/books", bookHandler.GetBooks)
	e.GET("/books/:id", bookHandler.GetBookByID)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)
}
