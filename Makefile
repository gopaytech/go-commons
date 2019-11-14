SHELL := /bin/bash

.EXPORT_ALL_VARIABLES:
OUT_DIR := ./_output
BIN_DIR := ./bin

$(shell mkdir -p $(OUT_DIR) $(BIN_DIR))

# Code build targets
.PHONY: vendor
vendor:
	go mod vendor

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./...

# Main Test Targets (without docker)
.PHONY: test
test:
	ENABLE_INTEGRATION_TEST=false \
	go test -race -coverprofile=$(OUT_DIR)/coverage.out ./...

.PHONY: itest
itest:
	ENABLE_INTEGRATION_TEST=true \
	go test -p 1 -race -coverprofile=$(OUT_DIR)/coverage.out ./...

