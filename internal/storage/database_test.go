package storage

import (
	"book-lib/internal/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testDB *gorm.DB

const (
	testUser     = "myuser"
	testPass     = "mypassword"
	testHost     = "localhost"
	testPort     = "3306"
	testDatabase = "test_booklib"
)

// setupTestDB создает подключение к тестовой БД и выполняет миграции
func setupTestDB(t *testing.T) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		testUser, testPass, testHost, testPort, testDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Ошибка подключения к тестовой БД: %v", err)
	}

	// Обновляем глобальную переменную testDB
	testDB = db

	// Удаляем и создаем таблицы заново
	err = db.Migrator().DropTable(&models.Book{})
	if err != nil {
		t.Fatalf("Ошибка удаления таблицы: %v", err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		t.Fatalf("Ошибка миграции: %v", err)
	}
}

// Очистка тестовой таблицы после каждого теста
func clearDB(t *testing.T) {
	err := testDB.Exec("DELETE FROM books").Error
	if err != nil {
		t.Fatalf("Ошибка очистки БД: %v", err)
	}
}

// Тест на добавление книги
func TestAddBook(t *testing.T) {
	setupTestDB(t)
	defer clearDB(t) // при необходимости можно закомментировать, чтобы посмотреть запись

	storage := NewBookStorage(testDB)
	book := models.Book{ID: 1, Title: "Тестовая книга", Author: "Тестовый Автор", Year: 2017}

	err := storage.AddBook(book)
	assert.NoError(t, err, "Ошибка при добавлении книги")
}

// Тест на получение списка книг
func TestGetBooks(t *testing.T) {
	setupTestDB(t)
	defer clearDB(t)

	storage := NewBookStorage(testDB)

	// Тестовая книга
	book := models.Book{ID: 1, Title: "Тестовая книга", Author: "Тестовый Автор", Year: 2017}
	err := storage.AddBook(book)
	assert.NoError(t, err)

	// Получение книг
	books, err := storage.GetBooks()
	assert.NoError(t, err, "Ошибка при получении книг")
	assert.Len(t, books, 1, "Ожидалась одна книга")
}

// Тест на получение книги по ID
func TestGetBookByID(t *testing.T) {
	setupTestDB(t)
	defer clearDB(t)

	storage := NewBookStorage(testDB)

	// Тестовая книга
	book := models.Book{ID: 1, Title: "Тестовая книга", Author: "Тестовый Автор", Year: 2017}
	err := storage.AddBook(book)
	assert.NoError(t, err)

	// Список книг, чтобы узнать ID
	books, err := storage.GetBooks()
	assert.NoError(t, err)
	assert.NotEmpty(t, books)

	// Получение книги по ID
	retrievedBook, err := storage.GetBookByID(books[0].ID)
	assert.NoError(t, err, "Ошибка при получении книги по ID")
	assert.Equal(t, "Тестовая книга", retrievedBook.Title)
	assert.Equal(t, "Тестовый Автор", retrievedBook.Author)
}

// Тест на обновление книги
func TestUpdateBook(t *testing.T) {
	setupTestDB(t)
	defer clearDB(t)

	storage := NewBookStorage(testDB)

	// Тестовая книга
	book := models.Book{ID: 1, Title: "Тестовая книга", Author: "Тестовый Автор", Year: 2017}
	err := storage.AddBook(book)
	assert.NoError(t, err)

	// ID добавленной книги
	books, err := storage.GetBooks()
	assert.NoError(t, err)
	assert.NotEmpty(t, books)
	bookID := books[0].ID

	// Обновление книги
	updatedBook := models.Book{Title: "Обновленное название", Author: "Обновленный автор", Year: 2044}
	err = storage.UpdateBook(bookID, updatedBook)
	assert.NoError(t, err, "Ошибка при обновлении книги")

	// Проверка обновленных данных
	retrievedBook, err := storage.GetBookByID(bookID)
	assert.NoError(t, err)
	assert.Equal(t, "Обновленное название", retrievedBook.Title)
	assert.Equal(t, "Обновленный автор", retrievedBook.Author)
	assert.Equal(t, 2044, retrievedBook.Year)
}

// Тест на удаление книги
func TestDeleteBook(t *testing.T) {
	setupTestDB(t)
	defer clearDB(t)

	storage := NewBookStorage(testDB)

	// Тестовая книга
	book := models.Book{ID: 1, Title: "Тестовая книга", Author: "Тестовый Автор", Year: 2017}
	err := storage.AddBook(book)
	assert.NoError(t, err)

	// ID добавленной книги
	books, err := storage.GetBooks()
	assert.NoError(t, err)
	assert.NotEmpty(t, books)
	bookID := books[0].ID

	// Удаление книги
	err = storage.DeleteBook(bookID)
	assert.NoError(t, err, "Ошибка при удалении книги...")

	// Проверка
	_, err = storage.GetBookByID(bookID)
	assert.Error(t, err, "Книга должна быть удалена...")
}
