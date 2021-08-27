export ROOT=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
export BIN=$(ROOT)/bin
export GOBIN?=$(BIN)
export GO=$(shell which go)
export BUILD=cd $(ROOT) && $(GO) install -v -ldflags "-s"
export CGO_ENABLED=0
export COMPOSE=$(shell which docker-compose)

all:
	$(BUILD) ./cmd/...

