fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golangci-lint run -v
.PHONY:lint

vet: lint
		go vet ./...
.PHONY:vet