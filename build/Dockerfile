FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o bin/dtservice ./cmd/dtservice/main.go

CMD ["./bin/dtservice"]