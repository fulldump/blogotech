

run:
	go run .

test:
	go test ./...

build:
	go build -o bin/blogotech .

deps:
	go mod download
