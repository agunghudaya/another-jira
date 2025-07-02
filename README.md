# 📊 Another Jira

A modern project management tool inspired by Jira, built with Go and React. This application provides a comprehensive solution for project tracking, task management, and team collaboration.

## 🏗️ Project Structure

```
another-jira/
├── be/                 # Backend service (Go)
├── fe/                 # Frontend application (React)
├── migrations/         # Database migrations
├── vault-data/        # Secure secret management
├── docker-compose.yml # Docker orchestration
└── README.md          # Project documentation
```

## 🚀 Key Features

- **Project Management**: Create and manage projects with customizable workflows
- **Task Tracking**: Track tasks, bugs, and stories with detailed information
- **Team Collaboration**: Assign tasks, add comments, and track progress
- **Real-time Updates**: Get instant notifications on task changes
- **Advanced Reporting**: Generate custom reports and analytics
- **User Management**: Role-based access control and team management

## ⚙️ Tech Stack

### Backend (Go)
- Clean Architecture for maintainable and scalable code
- PostgreSQL for data persistence
- RESTful API design
- JWT authentication
- Docker containerization

### Frontend (React)
- Modern React with functional components
- Material-UI for consistent design
- Redux for state management
- Responsive design for all devices
- Docker containerization

### Infrastructure
- Docker Compose for orchestration
- PostgreSQL database
- Vault for secret management
- Automated database migrations

## 🛠 Getting Started

### Prerequisites
- Docker and Docker Compose
- Node.js 16+ (for local frontend development)
- Go 1.18+ (for local backend development)
- PostgreSQL (if running without Docker)

### Quick Start with Docker
```bash
# Clone the repository
git clone https://github.com/your-org/another-jira.git
cd another-jira

# Start all services
docker-compose up -d

# Access the application
Frontend: http://localhost:3000
Backend API: http://localhost:8080
```

### Local Development
1. **Backend Setup**
   ```bash
   cd be
   go mod tidy
   go run cmd/main.go
   ```

2. **Frontend Setup**
   ```bash
   cd fe
   npm install
   npm start
   ```

3. **Database Setup**
   ```bash
   # Using Docker
   docker-compose up -d postgres
   
   # Or install PostgreSQL locally
   # Then run migrations
   cd migrations
   ./migrate.sh
   ```

## 📚 Documentation

- [Backend Documentation](be/README.md)
- [Frontend Documentation](fe/README.md)
- [API Documentation](be/docs/api.md)
- [Database Schema](migrations/README.md)

## 🔐 Security

- JWT-based authentication
- Role-based access control
- Secure password hashing
- Vault integration for secrets
- HTTPS support
- Input validation and sanitization

## 🧪 Testing

### Backend Tests
```bash
cd be
go test ./...
```

### Frontend Tests
```bash
cd fe
npm test
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by Atlassian Jira
- Built with modern best practices
- Community-driven development
