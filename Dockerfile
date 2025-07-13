# Start from the official Golang image for building
FROM golang:1.24.5 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o meme-api ./cmd/meme-api/main.go

# Use a minimal base image for running
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary from builder
COPY --from=builder /app/meme-api .

# Expose port (change if your app uses a different port)
EXPOSE 8080

# Command to run
CMD ["./meme-api"]
