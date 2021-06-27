#!make
include .bingo/Variables.mk

FILES_TO_FMT      ?= $(shell find . -name '*.go' -print)

# Ensure everything works even if GOPATH is not set, which is often the case.
# The `go env GOPATH` will work for all cases for Go 1.8+.
GOPATH      ?= $(shell go env GOPATH)
GOBIN       ?= $(firstword $(subst :, ,${GOPATH}))/bin
GOTEST_OPTS ?= --race -failfast -timeout 10m
GOPROXY     ?= https://proxy.golang.org

# Support gsed on OSX (installed via brew), falling back to sed. On Linux
# systems gsed won't be installed, so will use sed as expected.
SED     ?= $(shell which gsed 2>/dev/null || which sed)
GIT     ?= $(shell which git)

BIN_DIR ?= /tmp/bin
OS      ?= $(shell uname -s | tr '[A-Z]' '[a-z]')
ARCH    ?= $(shell uname -m)

SHELLCHECK ?= $(BIN_DIR)/shellcheck

define require_clean_work_tree
	@git update-index -q --ignore-submodules --refresh

	@if ! git diff-files --quiet --ignore-submodules --; then \
		echo >&2 "$1: you have unstaged changes."; \
		git diff-files --name-status -r --ignore-submodules -- >&2; \
		echo >&2 "Please commit or stash them."; \
		exit 1; \
	fi

	@if ! git diff-index --cached --quiet HEAD --ignore-submodules --; then \
		echo >&2 "$1: your index contains uncommitted changes."; \
		git diff-index --cached --name-status -r --ignore-submodules HEAD -- >&2; \
		echo >&2 "Please commit or stash them."; \
		exit 1; \
	fi

endef


.PHONY: deps
deps: ## Ensures fresh go.mod and go.sum.
	@go mod tidy
	@go mod verify

# The `validate` target checks for errors and inconsistencies in 
# our specification of an API. This target can check if we're 
# referencing inexistent definitions and gives us hints to where
# to fix problems with our API in a static manner.
validate:
	@swagger validate ./pkg/openapi/swagger.yml

# The `gen` target depends on the `validate` target as
# it will only succesfully generate the code if the specification
# is valid.
# 
# Here we're specifying some flags:
# --target              the base directory for generating the files;
# --spec                path to the swagger specification;
# --exclude-main        generates only the library code and not a 
#                       sample CLI application;
# --name                the name of the application.
.PHONY: gen
gen: validate 
	@swagger generate server \
		--target=./pkg/openapi/swagger \
		--spec=./pkg/openapi/swagger.yml \
		--exclude-main \
		--name=polydefi

.PHONY: check-git
check-git:
ifneq ($(GIT),)
	@test -x $(GIT) || (echo >&2 "No git executable binary found at $(GIT)."; exit 1)
else
	@echo >&2 "No git binary found."; exit 1
endif

.PHONY: build
build: ## Build the project.
build: check-git
build: export GIT_TAG=$(shell git describe --tags)
build: export GIT_HASH=$(shell git rev-parse --short HEAD)
build:
	@[ "${GIT_TAG}" ] || ( echo ">> GIT_TAG is not set"; exit 1 )
	@[ "${GIT_HASH}" ] || ( echo ">> GIT_HASH is not set"; exit 1 )
	go build -ldflags "-X main.GitTag=$(GIT_TAG) -X main.GitHash=$(GIT_HASH) -s -w" ./cmd



.PHONY: generate-bindings
generate-bindings: 
	rm -rf tmp/ioTube
	git clone https://github.com/iotexproject/ioTube tmp/ioTube
	rm -rf pkg/contracts/*
	mkdir pkg/contracts/tokenCashier
	abigen --sol tmp/ioTube/contracts/iotube/TokenCashier.sol --pkg tokenCashier    --out pkg/contracts/tokenCashier/tokenCashier.go
	mkdir pkg/contracts/tokenSafe
	abigen --sol tmp/ioTube/contracts/iotube/TokenSafe.sol --pkg tokenSafe    --out pkg/contracts/tokenSafe/tokenSafe.go
	mkdir pkg/contracts/tokenList
	abigen --sol tmp/ioTube/contracts/iotube/TokenList.sol --pkg tokenList    --out pkg/contracts/tokenList/tokenList.go
	mkdir pkg/contracts/shadowTokenList
	abigen --sol tmp/ioTube/contracts/iotube/ShadowTokenListManager.sol --pkg shadowTokenList    --out pkg/contracts/shadowTokenList/shadowTokenList.go
	# ERC20 token binding ->
	abigen --abi pkg/contracts/erc20/erc20.abi --pkg erc20    --out pkg/contracts/erc20/erc20.go