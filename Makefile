.DEFAULT_GOAL := build

fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golangci-lint run -v
.PHONY:lint

vet: fmt
		go vet ./...
.PHONY:vet

build: vet
		docker build -t url_shortener:local .
.PHONY:build

run: vet build
		docker compose up 
.PHONY:run