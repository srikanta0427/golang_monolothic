.PHONY: build run
build:
	go build -o bin/listing_app ./cmd/api
run: build
	./bin/listing_app
