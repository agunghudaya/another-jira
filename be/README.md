# 📊 Another Jira Backend

This is the backend service for the **Another Jira** project, a tool designed to synchronize Jira tasks, histories, and other data with a custom database. It enables advanced reporting, real-time task monitoring, and performance analysis.

---

## 🚀 Features

- **Jira Synchronization**: Sync Jira issues, histories, and user data.
- **PostgreSQL Integration**: Store Jira data in a relational database for advanced querying.
- **RESTful APIs**: Expose endpoints for interacting with Jira data.
- **Clean Architecture**: Ensures scalability, maintainability, and testability.
- **Logging**: Comprehensive logging using Logrus for better observability.
- **Extensibility**: Modular design for easy addition of new features.

---

## 🏗️ Architecture

The backend is built using **clean architecture** principles, ensuring a clear separation of concerns:

1. **Domain Layer**:
   - Core business logic, entities, and repository interfaces.
   - Independent of frameworks or external libraries.

2. **Usecase Layer**:
   - Application-specific workflows and orchestration.
   - Depends only on the domain layer.

3. **Infrastructure Layer**:
   - Handles external dependencies like databases, logging, and APIs.
   - Implements repository interfaces and other abstractions.

---

## 📂 Folder Structure

```
another-jira/
├── be/
│   ├── internal/
│   │   ├── domain/          # Core business logic and entities
│   │   ├── usecase/         # Application workflows
│   │   ├── infrastructure/  # External dependencies (e.g., database, logging)
│   │   └── repository/      # Repository implementations
│   ├── cmd/                 # Application entry points
│   └── config/              # Configuration files
└── ...
```

---

## ⚙️ Setup and Installation

### Prerequisites

- Go 1.18+
- PostgreSQL
- Docker (optional, for running the database locally)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/another-jira.git
   cd another-jira/be
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure the environment:
   - Copy the `.env.example` file to `.env` and update the values as needed.

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

5. (Optional) Run the database using Docker:
   ```bash
   docker-compose up -d
   ```

---

## 📖 API Documentation

The backend exposes the following APIs:

- **GET /users**: Fetch all Jira users.
- **POST /sync**: Trigger a Jira synchronization process.
- **GET /issues**: Retrieve Jira issues from the database.

For detailed API documentation, refer to the [API Docs](docs/api.md).

---

## 🧪 Testing

Run unit tests using the following command:

```bash
go test ./...
```

---

## 🤝 Contribution Guidelines

We welcome contributions! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a detailed description of your changes.

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
