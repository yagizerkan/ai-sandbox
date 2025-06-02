package services

import (
	"book-tracker/models"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	ErrBookNotFound    = errors.New("book not found")
	ErrInvalidStatus   = errors.New("invalid book status")
	ErrDuplicateBook   = errors.New("book with same title and author already exists")
)

// BookService defines the interface for book operations
type BookService interface {
	CreateBook(req models.CreateBookRequest) (*models.Book, error)
	GetBook(id string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	GetBooksByStatus(status models.BookStatus) ([]*models.Book, error)
	UpdateBook(id string, req models.UpdateBookRequest) (*models.Book, error)
	DeleteBook(id string) error
}

// InMemoryBookService implements BookService using in-memory storage
type InMemoryBookService struct {
	books map[string]*models.Book
	mutex sync.RWMutex
}

// NewInMemoryBookService creates a new instance of InMemoryBookService
func NewInMemoryBookService() *InMemoryBookService {
	return &InMemoryBookService{
		books: make(map[string]*models.Book),
	}
}

// CreateBook creates a new book in the library
func (s *InMemoryBookService) CreateBook(req models.CreateBookRequest) (*models.Book, error) {
	if !models.IsValidStatus(req.Status) {
		return nil, ErrInvalidStatus
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check for duplicate books (same title and author)
	for _, book := range s.books {
		if book.Title == req.Title && book.Author == req.Author {
			return nil, ErrDuplicateBook
		}
	}

	now := time.Now()
	book := &models.Book{
		ID:        uuid.New().String(),
		Title:     req.Title,
		Author:    req.Author,
		Status:    req.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.books[book.ID] = book
	return book, nil
}

// GetBook retrieves a book by its ID
func (s *InMemoryBookService) GetBook(id string) (*models.Book, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	book, exists := s.books[id]
	if !exists {
		return nil, ErrBookNotFound
	}

	return book, nil
}

// GetAllBooks retrieves all books in the library
func (s *InMemoryBookService) GetAllBooks() ([]*models.Book, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	books := make([]*models.Book, 0, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}

	return books, nil
}

// GetBooksByStatus retrieves all books with a specific status
func (s *InMemoryBookService) GetBooksByStatus(status models.BookStatus) ([]*models.Book, error) {
	if !models.IsValidStatus(status) {
		return nil, ErrInvalidStatus
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	books := make([]*models.Book, 0)
	for _, book := range s.books {
		if book.Status == status {
			books = append(books, book)
		}
	}

	return books, nil
}

// UpdateBook updates an existing book
func (s *InMemoryBookService) UpdateBook(id string, req models.UpdateBookRequest) (*models.Book, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	book, exists := s.books[id]
	if !exists {
		return nil, ErrBookNotFound
	}

	// Validate status if provided
	if req.Status != nil && !models.IsValidStatus(*req.Status) {
		return nil, ErrInvalidStatus
	}

	// Check for duplicate if title or author is being updated
	if req.Title != nil || req.Author != nil {
		newTitle := book.Title
		newAuthor := book.Author

		if req.Title != nil {
			newTitle = *req.Title
		}
		if req.Author != nil {
			newAuthor = *req.Author
		}

		for existingID, existingBook := range s.books {
			if existingID != id && existingBook.Title == newTitle && existingBook.Author == newAuthor {
				return nil, ErrDuplicateBook
			}
		}
	}

	// Update fields
	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.Author != nil {
		book.Author = *req.Author
	}
	if req.Status != nil {
		book.Status = *req.Status
	}

	book.UpdatedAt = time.Now()
	return book, nil
}

// DeleteBook removes a book from the library
func (s *InMemoryBookService) DeleteBook(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.books[id]; !exists {
		return ErrBookNotFound
	}

	delete(s.books, id)
	return nil
}