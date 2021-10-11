.EXPORT_ALL_VARIABLES:
OUT_DIR := ./_output
BIN_DIR := ./bin

PACKAGE=github.com/gopaytech/go-commons
CURRENT_DIR=$(shell pwd)

VERSION=$(shell cat ${CURRENT_DIR}/VERSION)
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell if [ -z "`git status --porcelain`" ]; then git describe --exact-match --tags HEAD 2>/dev/null; fi)

$(shell mkdir -p $(OUT_DIR) $(BIN_DIR))

# perform static compilation
STATIC_BUILD?=true

override LDFLAGS += \
  -X ${PACKAGE}.version=${VERSION} \
  -X ${PACKAGE}.buildDate=${BUILD_DATE} \
  -X ${PACKAGE}.gitCommit=${GIT_COMMIT}

ifeq (${STATIC_BUILD}, true)
override LDFLAGS += -extldflags "-static"
endif

ifneq (${GIT_TAG},)
IMAGE_TAG=${GIT_TAG}
LDFLAGS += -X ${PACKAGE}.gitTag=${GIT_TAG}
else
IMAGE_TAG?=$(GIT_COMMIT)
endif

# Code build targets
.PHONY: vendor
vendor:
	go mod vendor

# Main Test Targets (without docker)
.PHONY: test
test:
	go test -race -coverprofile=$(OUT_DIR)/coverage.out ./...

# Integration test executed by github workflow
.PHONY: integration-test
integration-test:
	go test -v -race -tags=integration -coverprofile=$(OUT_DIR)/coverage.out ./...

# Local Integration test only able to be executed on local with docker engine present
.PHONY: local-integration-test
local-integration-test:
	go test -v -race -tags=local,integration -coverprofile=$(OUT_DIR)/coverage.out ./...
