// Package models Author: Роман Манько [@speakerkiller]
package models

// Book структура книги
type Book struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
