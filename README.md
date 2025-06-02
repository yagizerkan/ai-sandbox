# 📚 Book Tracker API

A simple book tracking application built with Go and the Gin framework. Users can add books to their library and track their reading status.

## Features

- ✅ Add new books to your library
- ✅ Track reading status: "to be read", "currently reading", "read"
- ✅ Update book status
- ✅ Filter books by status
- ✅ Delete books from library
- ✅ Prevent duplicate books (same title and author)
- ✅ RESTful API with JSON responses
- ✅ Web interface for easy interaction
- ✅ CORS enabled for frontend integration

## Architecture

The application follows a clean architecture pattern with clear separation of concerns:

```
book-tracker/
├── models/          # Data models and types
├── services/        # Business logic layer
├── handlers/        # HTTP handlers (controllers)
├── static/          # Web interface files
├── main.go          # Application entry point
└── go.mod           # Go module dependencies
```

### Components

- **Models**: Define the Book struct and related types
- **Services**: Implement business logic with an interface-based approach
- **Handlers**: Handle HTTP requests and responses
- **In-Memory Storage**: Simple storage implementation (easily replaceable with database)

## API Endpoints

### Books

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/books` | Get all books (supports `?status=` filter) |
| `POST` | `/api/v1/books` | Create a new book |
| `GET` | `/api/v1/books/:id` | Get a specific book by ID |
| `PUT` | `/api/v1/books/:id` | Update a book |
| `DELETE` | `/api/v1/books/:id` | Delete a book |
| `GET` | `/api/v1/books/statuses` | Get available book statuses |

### Health Check

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check endpoint |

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd book-tracker
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

The server will start on port 12000. You can access:
- Web Interface: http://localhost:12000
- API: http://localhost:12000/api/v1
- Health Check: http://localhost:12000/health

## Usage Examples

### Using the Web Interface

1. Open http://localhost:12000 in your browser
2. Use the form to add new books
3. Filter books by status using the buttons
4. Update book status or delete books using the action buttons

### Using the API

#### Add a new book
```bash
curl -X POST http://localhost:12000/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "status": "to be read"
  }'
```

#### Get all books
```bash
curl http://localhost:12000/api/v1/books
```

#### Filter books by status
```bash
curl "http://localhost:12000/api/v1/books?status=currently%20reading"
```

#### Update book status
```bash
curl -X PUT http://localhost:12000/api/v1/books/{book-id} \
  -H "Content-Type: application/json" \
  -d '{"status": "read"}'
```

#### Delete a book
```bash
curl -X DELETE http://localhost:12000/api/v1/books/{book-id}
```

## Data Models

### Book
```go
type Book struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    Author    string    `json:"author"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Book Statuses
- `"to be read"` - Books you plan to read
- `"currently reading"` - Books you're actively reading
- `"read"` - Books you've finished reading

## API Response Examples

### Get All Books
```json
{
  "books": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "title": "Clean Code",
      "author": "Robert C. Martin",
      "status": "currently reading",
      "created_at": "2025-06-02T19:06:40.806540823Z",
      "updated_at": "2025-06-02T19:07:00.600507265Z"
    }
  ],
  "count": 1
}
```

### Create Book
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "status": "to be read",
  "created_at": "2025-06-02T19:06:40.806540823Z",
  "updated_at": "2025-06-02T19:06:40.806540823Z"
}
```

### Error Response
```json
{
  "error": "Book with same title and author already exists"
}
```

## Development

### Running Tests
```bash
go test ./...
```

### Running the Demo
```bash
go run example_demo.go
```

### Building for Production
```bash
go build -o book-tracker main.go
./book-tracker
```

## Configuration

The application uses the following default settings:
- **Port**: 12000
- **Host**: 0.0.0.0 (all interfaces)
- **CORS**: Enabled for all origins
- **Storage**: In-memory (data is lost on restart)

## Future Enhancements

- [ ] Database persistence (PostgreSQL, MySQL, SQLite)
- [ ] User authentication and authorization
- [ ] Book categories and tags
- [ ] Reading progress tracking
- [ ] Book recommendations
- [ ] Import/export functionality
- [ ] Search functionality
- [ ] Book cover images
- [ ] Reading statistics and analytics
- [ ] API rate limiting
- [ ] Logging and monitoring

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [Gin CORS](https://github.com/gin-contrib/cors) - CORS middleware
- [Google UUID](https://github.com/google/uuid) - UUID generation

## License

This project is open source and available under the [MIT License](LICENSE).

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Support

If you have any questions or need help, please open an issue on GitHub.
