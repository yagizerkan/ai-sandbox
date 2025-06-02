package models

import (
	"time"
)

// BookStatus represents the reading status of a book
type BookStatus string

const (
	StatusToBeRead        BookStatus = "to be read"
	StatusCurrentlyReading BookStatus = "currently reading"
	StatusRead            BookStatus = "read"
)

// Book represents a book in the user's library
type Book struct {
	ID          string     `json:"id"`
	Title       string     `json:"title" binding:"required"`
	Author      string     `json:"author" binding:"required"`
	Status      BookStatus `json:"status" binding:"required"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// IsValidStatus checks if the provided status is valid
func IsValidStatus(status BookStatus) bool {
	switch status {
	case StatusToBeRead, StatusCurrentlyReading, StatusRead:
		return true
	default:
		return false
	}
}

// CreateBookRequest represents the request payload for creating a book
type CreateBookRequest struct {
	Title  string     `json:"title" binding:"required"`
	Author string     `json:"author" binding:"required"`
	Status BookStatus `json:"status" binding:"required"`
}

// UpdateBookRequest represents the request payload for updating a book
type UpdateBookRequest struct {
	Title  *string     `json:"title,omitempty"`
	Author *string     `json:"author,omitempty"`
	Status *BookStatus `json:"status,omitempty"`
}