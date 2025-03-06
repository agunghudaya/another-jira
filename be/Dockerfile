
FROM golang:1.22.2 AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the binaries
RUN go build -o /app/http-api ./cmd/http/main.go
RUN go build -o /app/cron-worker ./cmd/worker/main.go

FROM golang:1.22.2 AS runtime

WORKDIR /app

# Copy compiled binaries
COPY --from=builder /app/http-api /app/cron-worker /app/

# Copy config files
COPY internal/infrastructure/config /app/internal/infrastructure/config

# Ensure executables have correct permissions
RUN chmod +x /app/http-api /app/cron-worker

# Default command (overridden in docker-compose)
CMD ["/app/http-api"]
