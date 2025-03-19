package handlers

import (
	"book-lib/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// MockBookService для тестирования
type MockBookService struct{}

func (m *MockBookService) GetBooks() ([]models.Book, error) {
	return []models.Book{{ID: 1, Title: "Test Book", Author: "Author"}}, nil
}
func (m *MockBookService) GetBookByID(id int64) (*models.Book, error) {
	if id == 1 {
		return &models.Book{ID: 1, Title: "Test Book", Author: "Author"}, nil
	}
	return &models.Book{}, errors.New("book not found")
}
func (m *MockBookService) AddBook(book models.Book) error              { return nil }
func (m *MockBookService) UpdateBook(id int64, book models.Book) error { return nil }
func (m *MockBookService) DeleteBook(id int64) error                   { return nil }

func setupHandler() *BookHandler {
	return NewBookHandler(&MockBookService{})
}

func TestGetBooks(t *testing.T) {
	e := echo.New()
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.GetBooks(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateBook(t *testing.T) {
	e := echo.New()
	h := setupHandler()
	book := models.Book{Title: "Test Book", Author: "Author Name"}
	body, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.CreateBook(c))
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetBookByID(t *testing.T) {
	e := echo.New()
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.GetBookByID(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateBook(t *testing.T) {
	e := echo.New()
	h := setupHandler()
	book := models.Book{Title: "Updated Title", Author: "New Author"}
	body, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPut, "/books/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.UpdateBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteBook(t *testing.T) {
	e := echo.New()
	h := setupHandler()
	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.DeleteBook(c))
	assert.Equal(t, http.StatusNoContent, rec.Code)
}
