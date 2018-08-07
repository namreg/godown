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

.PHONY: mocks
mocks:
	minimock -i github.com/namreg/godown-v2/pkg/clock.Clock -o ./pkg/clock -s "_mock.go" -b test
	minimock -i github.com/namreg/godown-v2/internal/pkg/command.Command -o ./internal/pkg/command -s "_mock.go" -b test
	minimock -i github.com/namreg/godown-v2/internal/pkg/storage.Storage -o ./internal/pkg/storage -s "_mock.go" -b test
	minimock -i net.Conn -o ./internal/pkg/server
