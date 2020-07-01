.DEFAULT_GOAL := build

GOFMT=gofmt
GC=go build
VERSION := $(shell git describe --abbrev=4 --dirty --always --tags)
Minversion := $(shell date)
IDENTIFIER= $(GOOS)-$(GOARCH)

help:          ## Show available options with this Makefile
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -v grep | awk 'BEGIN { FS = ":.*?##" }; { printf "%-15s  %s\n", $$1,$$2 }'

.PHONY: build
build: clean
	$(GC) -o ./bin/ethfs cmd/node/node.go

.PHONY: build
win: clean
	CGO_ENABLED=0  GOOS=windows GOARCH=amd64 $(GC)  -o ./bin/ethfs.exe cmd/node/node.go

.PHONY: glide
glide:   ## Installs glide for go package management
	@ mkdir -p $$(go env GOPATH)/bin
	@ curl https://glide.sh/get | sh;

vendor: glide.yaml glide.lock
	@ glide install

.PHONY: clean
clean:
	rm -rf ./bin/ethfs
	rm -rf build/
	rm -rf *.log

.PHONY: pb
pb:
	protoc -I=. -I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf --gogoslick_out=. pb/*.proto
