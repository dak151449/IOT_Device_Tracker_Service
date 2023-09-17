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

.PHONY: compose-up
compose-up:
	docker-compose build
	docker-compose up -d

.PHONY: compose-down
compose-down:
	docker-compose down