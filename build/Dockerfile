FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o bin/dt_service ./cmd/dt_service/main.go

CMD ["./bin/device-tracker-service"]