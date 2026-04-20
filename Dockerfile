# Build stage
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 ensures a static binary for minimal alpine image
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o hh-service ./cmd/app/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for connectivity (e.g. MongoDB Atlas)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/hh-service .

# Expose the application port
EXPOSE 8080

# Environment variables (can be overridden at runtime)
ENV GIN_MODE=release

# Start the application
CMD ["./hh-service"]
