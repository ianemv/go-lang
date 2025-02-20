# Dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /build

# Install required system dependencies
RUN apk add --no-cache gcc musl-dev git

# Install swag for Swagger docs generation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && \
    go mod verify && \
    go get github.com/gin-contrib/cors

# Copy the source code
COPY . .

# Generate Swagger docs
RUN swag init

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary and docs from builder
COPY --from=builder /build/main .
COPY --from=builder /build/docs ./docs

# Set executable permissions
RUN chmod +x /app/main

EXPOSE 8080

CMD ["/app/main"]
