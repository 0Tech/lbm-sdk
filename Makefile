#!/usr/bin/make -f

########################################
### Setup flags

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
# VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
VERSION :=v0.1.0
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true

export GO111MODULE = on

########################################
### Process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

########################################
### Process linker flags

ldflags = -X github.com/link-chain/link/version.Name=link \
		  -X github.com/link-chain/link/version.ServerName=linkd \
		  -X github.com/link-chain/link/version.ClientName=linkcli \
		  -X github.com/link-chain/link/version.Version=$(VERSION) \
		  -X github.com/link-chain/link/version.Commit=$(COMMIT) \
		  -X "github.com/link-chain/link/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'


########################################
### Build

all: install lint check

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/linkd.exe ./cmd/linkd
	go build -mod=readonly $(BUILD_FLAGS) -o build/linkcli.exe ./cmd/linkcli
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/linkd ./cmd/linkd
	go build -mod=readonly $(BUILD_FLAGS) -o build/linkcli ./cmd/linkcli
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

build-contract-tests-hooks:
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/contract_tests.exe ./cmd/contract_tests
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/contract_tests ./cmd/contract_tests
endif

install: go.sum
	go install $(BUILD_FLAGS) ./cmd/linkd
	go install $(BUILD_FLAGS) ./cmd/linkcli

install-debug: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/linkdebug



########################################
### Tools & dependencies

get-tools:
	go get github.com/rakyll/statik
	go get -u github.com/client9/misspell/cmd/misspell
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

clean:
	rm -rf  build/

########################################
### Testing


check: check-unit check-build
check-all: check check-race check-cover

check-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./...

check-race:
	@VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock' ./...

check-cover:
	@go test -mod=readonly -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock' ./...

check-build: build
	@go test -mod=readonly -p 4 `go list ./cli_test/...` -tags=cli_test -v


lint: golangci-lint
	golangci-lint run
	find . -name '*.go' -type f -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

format:
	find . -name '*.go' -type f -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs goimports -w -local github.com/link-chain/link

benchmark:
	@go test -mod=readonly -bench=. ./...


########################################
### Local validator nodes using docker and docker-compose

build-docker-testnet:
	$(MAKE) -C  $(CURDIR)/networks/local

build-conf-testnet:
	rm -rf $(CURDIR)/build/gentxs
	rm -rf $(CURDIR)/build/node*
	docker run --rm -v $(CURDIR)/build:/linkd:Z line/linkdnode testnet --v 4 -o . --starting-ip-address 192.168.10.2


# Run a 4-node testnet locally
start-testnet: stop-testnet
	docker-compose up -d

# Stop testnet
stop-testnet:
	docker-compose down

########################################
### Integration Test with multi containers

build-docker-integration:
	docker build --tag line/linkdnode-integtest .

check-build-integration:
	@go test -mod=readonly -p 4 `go list ./cli_test/...` -tags=cli_multi_node_test -v


########################################
### Simulation

# include simulations
include sims.mk


########################################
### Utilities

build-utils:
	$(MAKE) -C $(CURDIR)/contrib/cmd build

install-utils:
	$(MAKE) -C $(CURDIR)/contrib/cmd install

clean-utils:
	$(MAKE) -C $(CURDIR)/contrib/cmd clean


.PHONY: all build-linux install install-debug \
	go-mod-cache draw-deps clean build \
	build-utils install-utils clean-utils \
	check check-all check-build check-cover check-unit check-race

