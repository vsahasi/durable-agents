.PHONY: test vet fmt build

test: vet
	go test -race ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

build:
	go build -o bin/durable-agents ./cmd/durable-agents
