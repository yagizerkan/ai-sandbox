package handlers

import (
	"book-tracker/models"
	"book-tracker/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BookHandler handles HTTP requests for book operations
type BookHandler struct {
	bookService services.BookService
}

// NewBookHandler creates a new BookHandler
func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

// CreateBook handles POST /books
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req models.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.CreateBook(req)
	if err != nil {
		switch err {
		case services.ErrInvalidStatus:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book status. Must be 'to be read', 'currently reading', or 'read'"})
		case services.ErrDuplicateBook:
			c.JSON(http.StatusConflict, gin.H{"error": "Book with same title and author already exists"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		}
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBook handles GET /books/:id
func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")

	book, err := h.bookService.GetBook(id)
	if err != nil {
		if err == services.ErrBookNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book"})
		}
		return
	}

	c.JSON(http.StatusOK, book)
}

// GetAllBooks handles GET /books
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	status := c.Query("status")

	var books []*models.Book
	var err error

	if status != "" {
		books, err = h.bookService.GetBooksByStatus(models.BookStatus(status))
		if err != nil {
			if err == services.ErrInvalidStatus {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status parameter"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
			}
			return
		}
	} else {
		books, err = h.bookService.GetAllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"books": books, "count": len(books)})
}

// UpdateBook handles PUT /books/:id
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var req models.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.UpdateBook(id, req)
	if err != nil {
		switch err {
		case services.ErrBookNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		case services.ErrInvalidStatus:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book status. Must be 'to be read', 'currently reading', or 'read'"})
		case services.ErrDuplicateBook:
			c.JSON(http.StatusConflict, gin.H{"error": "Book with same title and author already exists"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		}
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook handles DELETE /books/:id
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := h.bookService.DeleteBook(id)
	if err != nil {
		if err == services.ErrBookNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// GetBookStatuses handles GET /books/statuses
func (h *BookHandler) GetBookStatuses(c *gin.Context) {
	statuses := []string{
		string(models.StatusToBeRead),
		string(models.StatusCurrentlyReading),
		string(models.StatusRead),
	}

	c.JSON(http.StatusOK, gin.H{"statuses": statuses})
}