
VERSION ?= 1.0.0
CURRDIR := $(shell pwd)

override BUILD_BIN           := $(CURRDIR)/bin

# ----------------------------------------------------------------------------------------------------------------------
# Set the location of various build tools
# ----------------------------------------------------------------------------------------------------------------------
override BUILD_OUTPUT        := $(CURRDIR)/build/_output
override BUILD_BIN           := $(CURRDIR)/bin
override BUILD_TARGETS       := $(BUILD_OUTPUT)/targets
override TEST_LOGS_DIR       := $(BUILD_OUTPUT)/test-logs
override COVERAGE_DIR        := $(BUILD_OUTPUT)/coverage

# ----------------------------------------------------------------------------------------------------------------------
# Set the location of various build tools
# ----------------------------------------------------------------------------------------------------------------------
TOOLS_DIRECTORY   = $(CURRDIR)/build/tools
TOOLS_BIN         = $(TOOLS_DIRECTORY)/bin

# ----------------------------------------------------------------------------------------------------------------------
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
# ----------------------------------------------------------------------------------------------------------------------
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# ----------------------------------------------------------------------------------------------------------------------
# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
# ----------------------------------------------------------------------------------------------------------------------
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

GITCOMMIT              ?= $(shell git rev-list -1 HEAD)
GITREPO                := https://github.com/tmiddlet2666/ghstats.git
SOURCE_DATE_EPOCH      := $(shell git show -s --format=format:%ct HEAD)
DATE_FMT               := "%Y-%m-%dT%H:%M:%SZ"
BUILD_DATE             := $(shell date -u -d "@$SOURCE_DATE_EPOCH" "+${DATE_FMT}" 2>/dev/null || date -u -r "${SOURCE_DATE_EPOCH}" "+${DATE_FMT}" 2>/dev/null || date -u "+${DATE_FMT}")
BUILD_USER             := $(shell whoami)

LDFLAGS          = -X main.Version=$(VERSION)$(MILESTONE) -X main.Commit=$(GITCOMMIT) -X main.Date=$(BUILD_DATE) -X main.Author=$(BUILD_USER)
GOS              = $(shell find . -type f -name "*.go" ! -name "*_test.go")

# ======================================================================================================================
# Makefile targets start here
# ======================================================================================================================

# ----------------------------------------------------------------------------------------------------------------------
# Display the Makefile help - this is a list of the targets with a description.
# This target MUST be the first target in the Makefile so that it is run when running make with no arguments
# ----------------------------------------------------------------------------------------------------------------------
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


# ======================================================================================================================
# Build targets
# ======================================================================================================================
##@ Build

# ----------------------------------------------------------------------------------------------------------------------
# Clean-up all of the build artifacts
# ----------------------------------------------------------------------------------------------------------------------
.PHONY: clean
clean: ## Cleans the build
	@echo "Cleaning Project"
	-rm -rf build/_output
	-rm -rf bin
	@mkdir -p $(TEST_LOGS_DIR)
	@mkdir -p $(COVERAGE_DIR)

# ----------------------------------------------------------------------------------------------------------------------
# Configure the build properties
# ----------------------------------------------------------------------------------------------------------------------
$(BUILD_PROPS):
	@echo "Creating build directories"
	@mkdir -p $(BUILD_OUTPUT)
	@mkdir -p $(BUILD_BIN)

# ----------------------------------------------------------------------------------------------------------------------
# Internal make step that builds the Coherence CLI for local platform
# ----------------------------------------------------------------------------------------------------------------------
.PHONY: ghstats
ghstats: $(BUILD_BIN)/ghstats-local   ## Build the binary for the local platform

$(BUILD_BIN)/ghstats-local: $(BUILD_PROPS) $(GOS)
	@echo "Building ghstats"
	CGO_ENABLED=0 GO111MODULE=on go build -trimpath -ldflags "$(LDFLAGS)" -o $(BUILD_BIN)/ghstats ./ghstats


# ----------------------------------------------------------------------------------------------------------------------
# Executes golangci-lint to perform various code review checks on the source.
# ----------------------------------------------------------------------------------------------------------------------
.PHONY: golangci
golangci: $(TOOLS_BIN)/golangci-lint ## Go code review
	$(TOOLS_BIN)/golangci-lint run -v --timeout=5m ./pkg/...


# ======================================================================================================================
# Test targets
# ======================================================================================================================
##@ Test

# ----------------------------------------------------------------------------------------------------------------------
# Executes the Go unit tests
# ----------------------------------------------------------------------------------------------------------------------
.PHONY: test-ghstats
test-ghstats: test-clean gotestsum $(BUILD_PROPS) ## Run the unit tests
	CGO_ENABLED=0 $(GOTESTSUM) --format testname --junitfile $(TEST_LOGS_DIR)/cohctl-test.xml \
	  -- $(GO_TEST_FLAGS) -v -coverprofile=$(COVERAGE_DIR)/cover-unit.out ./pkg/cmd/... ./pkg/utils/...
	go tool cover -html=$(COVERAGE_DIR)/cover-unit.out -o $(COVERAGE_DIR)/cover-unit.html

# ----------------------------------------------------------------------------------------------------------------------
# Find or download gotestsum
# ----------------------------------------------------------------------------------------------------------------------
.PHONY: gotestsum
GOTESTSUM = $(TOOLS_BIN)/gotestsum
gotestsum: ## Download gotestsum locally if necessary.
	$(call go-get-tool,$(GOTESTSUM),gotest.tools/gotestsum@v0.5.2)

# ----------------------------------------------------------------------------------------------------------------------
# Cleans the test cache
# ----------------------------------------------------------------------------------------------------------------------
.PHONY: test-clean
test-clean: gotestsum ## Clean the go test cache
	@echo "Cleaning test cache"
	go clean -testcache

# ----------------------------------------------------------------------------------------------------------------------
# Obtain the golangci-lint binary
# ----------------------------------------------------------------------------------------------------------------------
$(TOOLS_BIN)/golangci-lint:
	@mkdir -p $(TOOLS_BIN)
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(TOOLS_BIN) v1.43.0


# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2) into $(TOOLS_BIN)" ;\
GOBIN=$(TOOLS_BIN) go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
