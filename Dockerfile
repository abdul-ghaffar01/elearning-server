# ---- Build Stage ----
FROM golang:1.24-alpine AS builder

# Set workdir
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of your source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# ---- Run Stage ----
FROM alpine:latest

# Install necessary CA certs
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy schema and query files
COPY database/ /app/database

# Copy binary from builder stage
COPY --from=builder /app/server .

# Expose application port (adjust if needed)
EXPOSE 4406

# Start the application
CMD ["./server"]
