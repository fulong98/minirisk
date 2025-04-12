# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
# Copy the backend source files directly into the WORKDIR
COPY backend/ .

# Build the application
# Build the main package from the current directory
RUN CGO_ENABLED=0 GOOS=linux go build -o minirisk-api .

# Final stage - Use golang image to have 'go run' available for development
FROM golang:1.21-alpine

WORKDIR /app

# Note: Binary copy/chmod steps removed as we are using 'go run' with mounted source

# Create necessary directories
RUN mkdir -p /app/logs

# Expose port
EXPOSE 8080

# Set environment variables
ENV ENV=production
ENV PORT=8080

# Run the application using go run for development (compiles and runs on start)
# This works with the volume mount ./backend:/app
CMD ["go", "run", "main.go"]
