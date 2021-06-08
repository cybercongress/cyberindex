VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT  := $(shell git log -1 --format='%H')

export GO111MODULE = on

###############################################################################
###                                   All                                   ###
###############################################################################

all: build

###############################################################################
###                                Build flags                              ###
###############################################################################

LD_FLAGS = -X github.com/desmos-labs/juno/version.Version=$(VERSION) \
	-X github.com/desmos-labs/juno/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

###############################################################################
###                                  Build                                  ###
###############################################################################

build: go.sum
	@echo "building cyberindex binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/cyberindex ./cmd/cyberindex
.PHONY: build

###############################################################################
###                                 Install                                 ###
###############################################################################

install: go.sum
	@echo "installing cyberindex binary..."
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/cyberindex
.PHONY: install
