# Stage 1: Build the application
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o simple_api .

# Stage 2: Create a lightweight container
FROM alpine:latest
WORKDIR /app/

# Copy the binary from the builder stage
COPY --from=builder /app/simple_api /app/

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./simple_api"]