# Stage 1: Build the Go binary using the official Golang Alpine image (new version)
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a statically linked Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/main.go

# Stage 2: Create a lightweight final image using the latest Alpine Linux version
FROM alpine:3.19

# Set the working directory for the final image
WORKDIR /app

# Copy the binary from the builder stage into the final image
COPY --from=builder /app/app .

# Expose the port the application listens on
EXPOSE 8081

# Optional: Add a health check (adjust the URL if your health endpoint is different)
HEALTHCHECK --interval=30s --timeout=3s CMD wget -q --spider http://localhost:8081/health || exit 1

# Run the binary as the container entrypoint
ENTRYPOINT ["./app"]