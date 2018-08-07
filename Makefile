PACKAGE=github.com/namreg/godown-v2
SRC_PKGS=$(shell go list ./... | grep -v vendor)

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
	@go test -tags test $(SRC_PKGS) -cover