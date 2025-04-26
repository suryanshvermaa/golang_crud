# Golang CRUD Application

## Project Overview
This is a Go-based CRUD (Create, Read, Update, Delete) application that provides a RESTful API for managing student records. The application follows clean architecture principles and uses modern Go practices.

## Project Structure
```
.
├── cmd/
│   └── golang-crud/     # Main application entry point
├── internal/
│   ├── config/         # Configuration management
│   ├── http/           # HTTP handlers and middleware
│   └── types/          # Data types and models
├── storage/            # Data storage implementations
├── config/             # Configuration files
└── docs/              # Documentation
```

## Go Concepts Used

### 1. Project Organization
- **cmd/**: Contains the main application entry points
- **internal/**: Private application code
- **pkg/**: Public library code (if any)
- **config/**: Configuration files
- **storage/**: Data storage implementations

### 2. Go Modules
The project uses Go modules for dependency management:
- `go.mod`: Defines module dependencies
- `go.sum`: Contains cryptographic checksums of dependencies

### 3. Configuration Management
- Uses `cleanenv` package for configuration management
- Supports both YAML and environment variables
- Configuration includes:
  - Environment (production/development)
  - HTTP server address
  - Storage path

### 4. HTTP Server
- Uses standard `net/http` package
- Implements graceful shutdown
- Supports signal handling (SIGINT, SIGTERM)

### 5. API Endpoints

#### Health Check
- **Endpoint**: `GET /health`
- **Response**: "healthy✅"
- **Purpose**: Server health monitoring

#### Student Signup
- **Endpoint**: `POST /api/signup`
- **Request Body**: Student information
- **Response**: Success message
- **Status Codes**:
  - 201: Created
  - 400: Bad Request (validation errors)
  - 500: Internal Server Error

### 6. Error Handling
- Custom error responses
- Validation error handling using `go-playground/validator`
- Structured logging using `slog`

### 7. Data Validation
- Uses `go-playground/validator` for request validation
- Implements custom validation rules for student data

## Setup and Configuration

### Prerequisites
- Go 1.21 or later
- SQLite (for data storage)

### Environment Variables
- `CONFIG_PATH`: Path to configuration file
- `ENV`: Environment (production/development)

### Configuration File
The application requires a YAML configuration file with the following structure:
```yaml
env: production
storage_path: /path/to/storage
http_server:
  address: :8080
```

## Running the Application

1. Set up configuration:
   ```bash
   export CONFIG_PATH=/path/to/config.yaml
   ```

2. Run the application:
   ```bash
   go run cmd/golang-crud/main.go
   ```

## API Usage Examples

### Health Check
```bash
curl http://localhost:8080/health
```

### Student Signup
```bash
curl -X POST http://localhost:8080/api/signup \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

## Best Practices Implemented

1. **Clean Architecture**
   - Separation of concerns
   - Dependency injection
   - Interface-based design

2. **Error Handling**
   - Structured error responses
   - Proper error wrapping
   - Validation error handling

3. **Configuration Management**
   - Environment-based configuration
   - Secure configuration loading
   - Default values

4. **Logging**
   - Structured logging with `slog`
   - Contextual information
   - Different log levels

5. **Security**
   - Input validation
   - Secure configuration
   - Proper error handling

## Future Improvements

1. Add more CRUD operations:
   - GET /api/students
   - GET /api/students/{id}
   - PUT /api/students/{id}
   - DELETE /api/students/{id}

2. Implement authentication and authorization

3. Add database migrations

4. Implement rate limiting

5. Add API documentation using Swagger/OpenAPI

6. Add unit and integration tests

## Contributing
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License
[Add your license information here]
