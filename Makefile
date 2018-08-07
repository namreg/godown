BIN_PATH=./build
SERVER_BINARY=$(BIN_PATH)/server

vendor:
	$(GOPATH)/bin/dep ensure

run-server: build-server
	./$(SERVER_BINARY)

build-server: vendor
	mkdir -p $(BIN_PATH)
	rm -f $(SERVER_BINARY)
	go build -o $(SERVER_BINARY) ./cmd/server

build-server-linux: vendor
	mkdir -p $(BIN_PATH)
	rm -f $(SERVER_BINARY)
	env GOOS=linux GOARCH=amd64 go build -o $(SERVER_BINARY)-linux ./cmd/server

.PHONY: generate
generate:
	@go generate $(shell go list ./... | grep -v vendor)