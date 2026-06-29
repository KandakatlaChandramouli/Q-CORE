BINARY=qcore-cli

.PHONY: fmt vet lint test build clean run

fmt:
	goimports -w .
	go fmt ./...

vet:
	go vet ./...

lint:
	staticcheck ./...

test:
	go test -v ./...

build:
	go build -o build/$(BINARY) ./cmd/qcore-cli

run:
	go run ./cmd/qcore-cli

clean:
	rm -rf build
