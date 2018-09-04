PACKAGE=github.com/namreg/godown
PKGS=$(shell go list ./... | grep -v vendor | grep -v cmd)
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
	@go generate $(PKGS)

.PHONY: lint
lint:
	@echo "======> start linter"
	@docker run --rm -t -v $(GOPATH)/src/$(PACKAGE):/go/src/$(PACKAGE) -w /go/src/$(PACKAGE) roistat/golangci-lint -c .golang.ci.yaml

.PHONY: test
test: vendor generate
	@echo "======> running tests"
	@go test -race $(PKGS) -cover

.PHONY: clear
clear:
	@echo "======> clearing artifacts"
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
