# ---- Build Stage ----
FROM golang:1.25 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (to leverage caching)
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

# Set workdir
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Expose application port (change if your app uses a different port)
EXPOSE 4406

# Start the app
CMD ["./server"]
