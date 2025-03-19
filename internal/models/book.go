// Package models Author: Роман Манько [@speakerkiller]
package models

// Book структура книги
type Book struct {
	ID     int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
