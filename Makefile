PACKAGE=github.com/namreg/godown-v2
SRC_PKGS=$(shell go list ./... | grep -v vendor | grep -v cmd)
CURTIME=$(shell date +%Y-%m-%dT%T%z)
COMMIT=$(shell git rev-parse HEAD)
ARTIFACTS=./build

.PHONY: default
default: test

.PHONY: vendor
vendor:
	$(GOPATH)/bin/dep ensure

.PHONY: generate
generate:
	@go generate $(SRC_PKGS)

.PHONY: lint
lint:
	@docker run --rm -t -v $(GOPATH)/src/$(PACKAGE):/go/src/$(PACKAGE) -w /go/src/$(PACKAGE) roistat/golangci-lint -c .golang.ci.yaml

.PHONY: test
test: vendor generate
	@go test -race $(SRC_PKGS) -cover

.PHONY: clear
clear:
	rm -fR $(ARTIFACTS)

.PHONY: build
build: clear
	go build -ldflags="-X main.buildtime=$(CURTIME) -X main.commit=$(COMMIT)" -o $(ARTIFACTS)/godown-server ./cmd/godown-server
	go build -ldflags="-X main.buildtime=$(CURTIME) -X main.commit=$(COMMIT)" -o $(ARTIFACTS)/godown-cli ./cmd/godown-cli