# Use Go as the builder
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the binaries
RUN go build -o /app/http-api ./cmd/http/main.go
RUN go build -o /app/cron-worker ./cmd/worker/main.go

# Use a minimal runtime image
FROM debian:bookworm-slim

WORKDIR /app

# Copy compiled binaries
COPY --from=builder /app/http-api /app/cron-worker /app/

# Copy config files to /app/internal/configs
COPY internal/configs /app/internal/configs

# Ensure executables have correct permissions
RUN chmod +x /app/http-api /app/cron-worker

# Default command (overridden in docker-compose)
CMD ["/app/http-api"]
