# Stage 1: Build the application
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o mocking_api .

# Stage 2: Create a container
FROM golang:1.21
WORKDIR /app/
COPY --from=builder /app/mocking_api /app/
CMD ["./mocking_api"]
