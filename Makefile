-include .env
export

ifeq ($(POSTGRES_SETUP),)
	POSTGRES_SETUP := "user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=localhost port=$(POSTGRES_PORT) sslmode=disable"
endif

.PHONY: migration-create migrate-up migrate_down
migration-create:
	goose -dir $(MIGRATION_FOLDER) create "$(name)" sql

migrate-up:
	goose -dir $(MIGRATION_FOLDER) postgres $(POSTGRES_SETUP) up

migrate-down:
	goose -dir $(MIGRATION_FOLDER) postgres $(POSTGRES_SETUP) down


.PHONY: generate
generate:
	mkdir -p pkg
	protoc  --go_out=./pkg --go_opt=paths=source_relative \
			--go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
 			api/device-tracker.proto

.PHONY: lint
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
	golangci-lint run ./...

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

.PHONY: run-db
run-db:
	docker-compose up -d postgres

.PHONY: compose-stop
compose-stop:
	docker-compose stop

.PHONY: compose-down
compose-down:
	docker-compose down

.PHONY: run-app
run-app:
	GOOS=linux GOARCH=amd64 $(MAKE) build
	$(MAKE) compose-up