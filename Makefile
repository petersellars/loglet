.PHONY: build run
.DEFAULT_GOAL := run

build:
	go build ./cmd/loglet

run:
	go run ./cmd/loglet --help

adr-readme:
	adrs generate toc -i ./docs/adr/templates/intro.md -o ./docs/adr/templates/outro.md > ./docs/adr/README.md
