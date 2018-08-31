PACKAGE=github.com/namreg/godown-v2
SRC_PKGS=$(shell go list ./... | grep -v vendor | grep -v cmd)
CURTIME=$(shell date +%Y-%m-%dT%T%z)
COMMIT=$(shell git rev-parse HEAD)
ARTIFACTS=./build

.PHONY: default
default: test

.PHONY: vendor
vendor:
	@echo "======> vendoring dependencies"
	@$(GOPATH)/bin/dep ensure

.PHONY: generate
generate:
	@echo "======> generating code"
	@go generate $(SRC_PKGS)

.PHONY: lint
lint:
	@echo "======> start linter"
	@docker run --rm -t -v $(GOPATH)/src/$(PACKAGE):/go/src/$(PACKAGE) -w /go/src/$(PACKAGE) roistat/golangci-lint -c .golang.ci.yaml

.PHONY: test
test: vendor generate
	@echo "======> running tests"
	@go test -race $(SRC_PKGS) -cover

.PHONY: clear
clear:
	@echo "======> clear artifacts"
	@rm -fR $(ARTIFACTS)

.PHONY: build
build: clear
	go build -o $(ARTIFACTS)/godown-server ./cmd/godown-server
	go build -o $(ARTIFACTS)/godown-cli ./cmd/godown-cli

.PHONY: release
release: test lint
	@echo "======> starting a new release"
	docker run --rm --privileged \
	-v $(PWD):/go/src/$(PACKAGE) \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-w /go/src/$(PACKAGE) \
	-e GITHUB_TOKEN \
	-e DOCKER_USERNAME \
	-e DOCKER_PASSWORD \
	goreleaser/goreleaser release --rm-dist
