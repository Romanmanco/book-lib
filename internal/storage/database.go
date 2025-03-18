package storage

import (
	"book-lib/config"
	"book-lib/internal/models"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB инициализирует подключение к базе данных
func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("DB_USER", "myuser"),
		config.GetEnv("DB_PASSWORD", "mypassword"),
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_PORT", "3306"),
		config.GetEnv("DB_NAME", "booklib"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Ошибка подключения к БД:", err)
		return nil, err
	}

	log.Info("Успешное подключение к БД")
	return db, nil
}

// RunMigrations миграции таблиц
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Error("Ошибка миграции:", err)
		return
	}

	log.Info("Миграции выполнены успешно...")
}

// BookStore интерфейс хранилища книг
type BookStore interface {
	AddBook(book models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	UpdateBook(id string, updatedBook models.Book) error
	DeleteBook(id string) error
}

// BookStorage структура для хранения книг в памяти
type BookStorage struct {
	db *gorm.DB
}

// NewBookStorage для создания нового хранилища
func NewBookStorage(db *gorm.DB) *BookStorage {
	return &BookStorage{db: db}
}

// AddBook добавить книгу
func (s *BookStorage) AddBook(book models.Book) error {
	if err := s.db.Create(&book).Error; err != nil {
		return fmt.Errorf("failed to add book: %w", err)
	}
	return nil
}

// GetBooks получить все книги
func (s *BookStorage) GetBooks() ([]models.Book, error) {
	var books []models.Book
	if err := s.db.Find(&books).Error; err != nil {
		return nil, fmt.Errorf("failed to get books: %w", err)
	}
	return books, nil
}

// GetBookByID получит книгу по ID
func (s *BookStorage) GetBookByID(id string) (*models.Book, error) {
	var book models.Book
	if err := s.db.First(&book, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, fmt.Errorf("failed to get book by ID: %w", err)
	}
	return &book, nil
}

// UpdateBook обновляет данные книги по ID
func (s *BookStorage) UpdateBook(id string, updatedBook models.Book) error {
	var book models.Book
	if err := s.db.First(&book, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("book not found")
		}
		return fmt.Errorf("failed to get book for update: %w", err)
	}
	if err := s.db.Model(&book).Updates(updatedBook).Error; err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}
	return nil
}

// DeleteBook удаление книги по ID
func (s *BookStorage) DeleteBook(id string) error {
	if err := s.db.Delete(&models.Book{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}
	return nil
}
