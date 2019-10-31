

run:
	go run .

test:
	go test -p=32 ./...

build:
	go build -o bin/blogotech .

deps:
	go mod download
