.PHONY: build run clean test

BINARY := clipwipe

build:
	go build -o $(BINARY)

run: build
	./$(BINARY)

run-debug: build
	./$(BINARY) -debug -interval 100ms

clean:
	rm -f $(BINARY)
	go clean

test:
	go test -v ./...

tidy:
	go mod tidy

fmt:
	go fmt ./...

lint:
	go vet ./...
