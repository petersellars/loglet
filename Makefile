.PHONY: build run
.DEFAULT_GOAL := run

build:
	go build ./cmd/loglet

run:
	go run ./cmd/loglet --help