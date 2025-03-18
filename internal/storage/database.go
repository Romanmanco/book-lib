// Package storage Author: Роман Манько [@speakerkiller]
package storage

import (
	"book-lib/internal/models"
	"errors"
	"sync"
)

// BookStore интерфейс хранилища книг
type BookStore interface {
	AddBook(book models.Book) error
	GetBooks() []models.Book
	GetBookByID(id string) (*models.Book, error)
	UpdateBook(id string, updatedBook models.Book) error
	DeleteBook(id string) error
}

// BookStorage структура для хранения книг в памяти
type BookStorage struct {
	mu    sync.Mutex
	books map[string]models.Book
}

// NewBookStorage для создания нового хранилища
func NewBookStorage() *BookStorage {
	return &BookStorage{
		books: make(map[string]models.Book),
	}
}

// AddBook добавить книгу
func (s *BookStorage) AddBook(book models.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// существуют ли книги с таким же ID
	if _, exists := s.books[book.ID]; exists {
		return errors.New("book with this ID already exists")
	}
	s.books[book.ID] = book
	return nil
}

// GetBooks получить все книги
func (s *BookStorage) GetBooks() []models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()

	books := make([]models.Book, 0, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}
	return books
}

// GetBookByID получит книгу по ID
func (s *BookStorage) GetBookByID(id string) (*models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	book, exists := s.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return &book, nil
}

// UpdateBook обновляет данные книги по ID
func (s *BookStorage) UpdateBook(id string, updatedBook models.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return errors.New("book not found")
	}
	s.books[id] = updatedBook
	return nil
}

// DeleteBook удаление книги по ID
func (s *BookStorage) DeleteBook(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(s.books, id)
	return nil
}
