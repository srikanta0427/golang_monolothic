.PHONY: build run
build:
	go build -o build/listing_app ./cmd/api
run:
	./build/listing_app
