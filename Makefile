-include .env
export

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: generate
generate:
	mkdir -p pkg
	protoc  --go_out=./pkg --go_opt=paths=source_relative \
			--go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
 			api/device-tracker.proto

.PHONY: build
build:
	go mod download
	go build -o bin/dtservice ./cmd/dtservice/main.go

.PHONY: run
run:
	go mod download
	go run ./cmd/dtservice/main.go

.PHONY: compose-up
compose-up:
	docker-compose up -d --build

.PHONY: compose-stop
compose-stop:
	docker-compose stop

.PHONY: compose-down
compose-down:
	docker-compose down

.PHONY: run-docker
run-docker:
	GOOS=linux GOARCH=amd64 $(MAKE) build
	$(MAKE) compose-up