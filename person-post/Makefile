.PHONY: test build run

test:
	go test -v ./...

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/v1 *.go

run:
	go run main.go
