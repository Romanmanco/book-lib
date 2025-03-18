// Package handlers Author: Роман Манько [@speakerkiller]
package handlers

import (
	"book-lib/internal/models"
	"book-lib/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

// BookHandler структура для хендлера
type BookHandler struct {
	service service.BookService
}

// NewBookHandler конструктор хендлера
func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{
		service: s,
	}
}

// CreateBook создание новой книги
func (h *BookHandler) CreateBook(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.service.AddBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, book)
}

// GetBooks получение списка книг
func (h *BookHandler) GetBooks(c echo.Context) error {
	books := h.service.GetBooks()
	return c.JSON(http.StatusOK, books)
}

// GetBookByID получение книги по ID
func (h *BookHandler) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	book, err := h.service.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

// UpdateBook обновление данных книги
func (h *BookHandler) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.service.UpdateBook(id, book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

// DeleteBook удаление книги
func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteBook(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
