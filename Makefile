.PHONY: build run

build:
	go build -o bin/fucit cmd/fucit/main.go

run: build
	bin/fucit