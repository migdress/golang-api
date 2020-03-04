.PHONY: test build run

test:
	go test -v ./...

build:
	export GO111MODULE=on
	env CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/v1 *.go 

run:
	go run main.go
