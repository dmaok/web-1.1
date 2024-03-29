.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: debug
debug:
	go run ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build