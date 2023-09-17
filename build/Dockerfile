FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o bin/device-tracker-service ./cmd/device-tracker-service/main.go

CMD ["./bin/device-tracker-service"]