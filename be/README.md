# 📊 Another Jira Backend

A robust backend service built with Go, implementing clean architecture principles for a modern project management system.

## 🏗️ Architecture

The backend follows clean architecture principles with clear separation of concerns:

```
be/
├── cmd/                    # Application entry points
│   └── main.go            # Main application entry
├── internal/              # Private application code
│   ├── domain/           # Core business logic and entities
│   │   ├── entity/       # Business entities
│   │   └── repository/   # Repository interfaces
│   ├── usecase/          # Application business rules
│   ├── delivery/         # Delivery mechanisms (HTTP, gRPC)
│   ├── repository/       # Repository implementations
│   ├── infrastructure/   # External services integration
│   └── middleware/       # Cross-cutting concerns
└── migrations/           # Database migrations
```

## 🚀 Features

- **RESTful API**: Well-documented HTTP endpoints
- **Authentication**: JWT-based authentication
- **Authorization**: Role-based access control
- **Database**: PostgreSQL with migrations
- **Logging**: Structured logging with Logrus
- **Error Handling**: Consistent error responses
- **Validation**: Request validation
- **Documentation**: Swagger/OpenAPI support

## ⚙️ Tech Stack

- **Language**: Go 1.18+
- **Framework**: Standard library + custom middleware
- **Database**: PostgreSQL
- **ORM**: Custom SQL with prepared statements
- **Authentication**: JWT
- **Logging**: Logrus
- **Testing**: Go testing package
- **Container**: Docker

## 🛠 Setup and Installation

### Prerequisites

- Go 1.18+
- PostgreSQL 13+
- Docker (optional)
- Make (optional, for using Makefile)

### Local Development

1. **Clone and Setup**
   ```bash
   git clone https://github.com/your-org/another-jira.git
   cd another-jira/be
   go mod tidy
   ```

2. **Environment Configuration**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Database Setup**
   ```bash
   # Using Docker
   docker-compose up -d postgres
   
   # Or connect to your PostgreSQL instance
   # Then run migrations
   ./scripts/migrate.sh
   ```

4. **Run the Application**
   ```bash
   go run cmd/main.go
   ```

### Docker Deployment

```bash
# Build the image
docker build -t another-jira-backend .

# Run the container
docker run -p 8080:8080 another-jira-backend
```

## 📚 API Documentation

The API follows RESTful principles and is documented using Swagger/OpenAPI.

### Key Endpoints

- `POST /api/v1/auth/login` - User authentication
- `GET /api/v1/projects` - List projects
- `POST /api/v1/projects` - Create project
- `GET /api/v1/issues` - List issues
- `POST /api/v1/issues` - Create issue
- `PUT /api/v1/issues/:id` - Update issue
- `GET /api/v1/users` - List users

For detailed API documentation, visit `/swagger/index.html` when running the server.

## 🧪 Testing

### Unit Tests
```bash
go test ./...
```

### Integration Tests
```bash
go test -tags=integration ./...
```

### Test Coverage
```bash
go test -cover ./...
```

## 🔐 Security

- JWT-based authentication
- Password hashing with bcrypt
- Input validation and sanitization
- CORS configuration
- Rate limiting
- SQL injection prevention
- XSS protection

## 📦 Dependencies

Key dependencies:
- `github.com/gin-gonic/gin` - HTTP framework
- `github.com/golang-jwt/jwt` - JWT implementation
- `github.com/lib/pq` - PostgreSQL driver
- `github.com/sirupsen/logrus` - Logging
- `github.com/spf13/viper` - Configuration

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.
