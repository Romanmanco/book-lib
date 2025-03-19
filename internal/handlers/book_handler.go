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
// @Summary Create a new book
// @Description Create a new book by providing the book details
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book details"
// @Success 201 {object} models.Book
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /books [post]
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
// @Summary Get all books
// @Description Get a list of all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Failure 404 {string} string "Books not found"
// @Router /books [get]
func (h *BookHandler) GetBooks(c echo.Context) error {
	books, err := h.service.GetBooks()
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

// GetBookByID получение книги по ID
// @Summary Get a book by ID
// @Description Get a single book by its ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {string} string "Book not found"
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	book, err := h.service.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

// UpdateBook обновление данных книги
// @Summary Update a book
// @Description Update a book by providing the updated details
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.Book true "Updated book details"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Book not found"
// @Failure 500 {string} string "Internal server error"
// @Router /books/{id} [put]
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
// @Summary Delete a book by ID
// @Description Delete a book by its ID
// @Tags books
// @Param id path string true "Book ID"
// @Success 204 "Book deleted successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteBook(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
