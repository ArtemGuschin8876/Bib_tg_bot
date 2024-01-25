.PHONY: build
build:
  go build

.PHONY: test
test:
  go test -v -race -timeout 30s ./ ...

.DEFAULT_GOAL := build
