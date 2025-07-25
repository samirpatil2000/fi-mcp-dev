FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Final stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/app .

# Copy static files and test data
COPY --from=builder /app/static ./static
COPY --from=builder /app/test_data_dir ./test_data_dir

# Copy .env file if it exists
COPY --from=builder /app/.env* ./

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./app"]