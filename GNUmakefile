TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
PKG_NAME=cfssl

default: build

build:
	go install

test:
	go test $(TEST) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v -parallel 20 $(TESTARGS) -timeout 120m

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)

.PHONY: build test testacc fmt
