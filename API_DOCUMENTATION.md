# Book Tracker API Documentation

A simple REST API for tracking books in your personal library. Built with Go and Gin framework.

## Features

- Add new books to your library
- Mark books with reading status: "to be read", "currently reading", or "read"
- Update book information and status
- Delete books from your library
- Filter books by status
- Prevent duplicate books (same title and author)

## API Endpoints

### Base URL
```
http://localhost:12000/api/v1
```

### Health Check
```
GET /health
```
Returns server health status.

### Book Endpoints

#### 1. Create a Book
```
POST /api/v1/books
```

**Request Body:**
```json
{
  "title": "The Go Programming Language",
  "author": "Alan Donovan",
  "status": "to be read"
}
```

**Response (201 Created):**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "title": "The Go Programming Language",
  "author": "Alan Donovan",
  "status": "to be read",
  "created_at": "2025-06-02T19:00:00Z",
  "updated_at": "2025-06-02T19:00:00Z"
}
```

#### 2. Get All Books
```
GET /api/v1/books
```

**Optional Query Parameters:**
- `status`: Filter by reading status ("to be read", "currently reading", "read")

**Response (200 OK):**
```json
{
  "books": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "title": "The Go Programming Language",
      "author": "Alan Donovan",
      "status": "to be read",
      "created_at": "2025-06-02T19:00:00Z",
      "updated_at": "2025-06-02T19:00:00Z"
    }
  ],
  "count": 1
}
```

#### 3. Get a Specific Book
```
GET /api/v1/books/{id}
```

**Response (200 OK):**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "title": "The Go Programming Language",
  "author": "Alan Donovan",
  "status": "to be read",
  "created_at": "2025-06-02T19:00:00Z",
  "updated_at": "2025-06-02T19:00:00Z"
}
```

#### 4. Update a Book
```
PUT /api/v1/books/{id}
```

**Request Body (all fields optional):**
```json
{
  "title": "Updated Title",
  "author": "Updated Author",
  "status": "currently reading"
}
```

**Response (200 OK):**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "title": "Updated Title",
  "author": "Updated Author",
  "status": "currently reading",
  "created_at": "2025-06-02T19:00:00Z",
  "updated_at": "2025-06-02T19:05:00Z"
}
```

#### 5. Delete a Book
```
DELETE /api/v1/books/{id}
```

**Response (200 OK):**
```json
{
  "message": "Book deleted successfully"
}
```

#### 6. Get Available Statuses
```
GET /api/v1/books/statuses
```

**Response (200 OK):**
```json
{
  "statuses": [
    "to be read",
    "currently reading",
    "read"
  ]
}
```

## Valid Book Statuses

- `"to be read"` - Books you plan to read
- `"currently reading"` - Books you are currently reading
- `"read"` - Books you have finished reading

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid book status. Must be 'to be read', 'currently reading', or 'read'"
}
```

### 404 Not Found
```json
{
  "error": "Book not found"
}
```

### 409 Conflict
```json
{
  "error": "Book with same title and author already exists"
}
```

### 500 Internal Server Error
```json
{
  "error": "Failed to create book"
}
```

## Example Usage

### Add a new book
```bash
curl -X POST http://localhost:12000/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "status": "to be read"
  }'
```

### Get all books currently being read
```bash
curl "http://localhost:12000/api/v1/books?status=currently reading"
```

### Update book status to read
```bash
curl -X PUT http://localhost:12000/api/v1/books/{book-id} \
  -H "Content-Type: application/json" \
  -d '{
    "status": "read"
  }'
```

## Running the Application

1. Install Go 1.21 or later
2. Clone the repository
3. Run `go mod tidy` to install dependencies
4. Run `go run main.go`
5. The API will be available at `http://localhost:12000`

## Architecture

The application follows a clean architecture pattern:

- **Models**: Data structures and validation
- **Services**: Business logic and data operations
- **Handlers**: HTTP request/response handling
- **Main**: Application setup and routing

The current implementation uses in-memory storage for simplicity, but the service interface makes it easy to swap in a database implementation later.