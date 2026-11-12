.PHONY: build run migrate-up migrate-down migrate-force

build:
	@go build -o bin/listing_app ./cmd/api

run: build
	@./bin/listing_app

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down