.PHONY: act-build adr-readme build run
.DEFAULT_GOAL := run

act-build:
	act -j build -P ubuntu-24.04=ghcr.io/catthehacker/ubuntu:act-24.04

adr-readme:
	adrs generate toc -i ./docs/adr/templates/intro.md -o ./docs/adr/templates/outro.md > ./docs/adr/README.md

build:
	go build ./cmd/loglet

run:
	go run ./cmd/loglet --help
