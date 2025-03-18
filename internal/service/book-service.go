// Package service Author: Роман Манько [@speakerkiller]
package service

import (
	"book-lib/internal/models"
	"book-lib/internal/storage"
)

// BookService интерфейс для работы с книгами
type BookService interface {
	AddBook(book models.Book) error
	GetBooks() []models.Book
	GetBookByID(id string) (*models.Book, error)
}

// bookService реализация BookService
type bookService struct {
	store storage.BookStore
}

// NewBookService конструктор сервиса для работы с книгами
func NewBookService(store storage.BookStore) BookService {
	return &bookService{
		store: store,
	}
}

// AddBook добавить книгу
func (s *bookService) AddBook(book models.Book) error {
	return s.store.AddBook(book)
}

// GetBooks все книги
func (s *bookService) GetBooks() []models.Book {
	return s.store.GetBooks()
}

// GetBookByID получить книгу по ID
func (s *bookService) GetBookByID(id string) (*models.Book, error) {
	return s.store.GetBookByID(id)
}
