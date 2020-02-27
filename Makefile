BINNAME   ?= windwaker

BINDIR    := $(CURDIR)/bin
GOPATH     = $(shell go env GOPATH)
DEP        = $(GOPATH)/bin/dep
GOIMPORTS  = $(GOPATH)/bin/goimports
ARCH       = $(shell uname -p)

LDFLAGS   := -w -s
GOFLAGS   :=
SRC       := $(shell find . -type f -name '*.go' -print)

# Required for globs to work properly
SHELL      = /usr/bin/env bash

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
TAGS       =

ifdef VERSION
	BINARY_VERSION = $(VERSION)
endif
BINARY_VERSION ?= ${GIT_TAG}

ifneq ($(BINARY_VERSION),)
	LDFLAGS += ''
endif

.PHONY: all
all: build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GOOS=linux GO111MODULE=on go build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(BINNAME) .

.PHONY: clean
clean:
	@rm -rf $(BINDIR)

.PHONY: info
info:
	@echo "Version:          ${VERSION}"
	@echo "Git Tag:          ${GIT_TAG}"
	@echo "Git Commit:       ${GIT_COMMIT}"
