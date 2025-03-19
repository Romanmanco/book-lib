package service

import (
	"book-lib/internal/models"
	"book-lib/internal/storage"
)

// BookService интерфейс для работы с книгами
type BookService interface {
	AddBook(book models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookByID(id int64) (*models.Book, error)
	UpdateBook(id int64, updatedBook models.Book) error
	DeleteBook(id int64) error
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
func (s *bookService) GetBooks() ([]models.Book, error) {
	return s.store.GetBooks()
}

// GetBookByID получить книгу по ID
func (s *bookService) GetBookByID(id int64) (*models.Book, error) {
	return s.store.GetBookByID(id)
}

// UpdateBook обновляет книгу
func (s *bookService) UpdateBook(id int64, updatedBook models.Book) error {
	return s.store.UpdateBook(id, updatedBook)
}

// DeleteBook удаляет книгу
func (s *bookService) DeleteBook(id int64) error {
	return s.store.DeleteBook(id)
}
