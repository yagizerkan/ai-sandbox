package main

import (
	"book-tracker/handlers"
	"book-tracker/services"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize services
	bookService := services.NewInMemoryBookService()

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookService)

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	// Serve static files
	router.Static("/static", "./static")
	router.StaticFile("/", "./static/index.html")

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Book routes
		books := api.Group("/books")
		{
			books.POST("", bookHandler.CreateBook)
			books.GET("", bookHandler.GetAllBooks)
			books.GET("/statuses", bookHandler.GetBookStatuses)
			books.GET("/:id", bookHandler.GetBook)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}
	}

	// Start server
	port := ":12000"
	log.Printf("Starting Book Tracker API server on port %s", port)
	log.Printf("Health check: http://localhost%s/health", port)
	log.Printf("API documentation available at: http://localhost%s/api/v1/books/statuses", port)
	
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}