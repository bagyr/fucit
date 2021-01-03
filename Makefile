.PHONY: build run test

build:
	go build -o bin/fucit cmd/fucit/main.go

run: build
	bin/fucit

test: build
	go test -v ./...