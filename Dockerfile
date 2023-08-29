FROM golang:1.17
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o ./mocking_api .
CMD ./mocking_api