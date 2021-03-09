MAINDIR := $(shell find . -type f -name 'main.go' -printf '%h\n')
APPNAME := $(shell basename $(MAINDIR))
TARGET := $(shell echo ${PWD}/bin/$(APPNAME))

VERSION := $(shell git describe --tags --abbrev=0)
BUILD := $(shell git rev-list -1 $(VERSION))
LDFLAGS := -ldflags "-X=main.Version=$(VERSION) -X=main.Appname=$(APPNAME)"


.PHONY: install build run clean uninstall test check

install: check
	@go install $(LDFLAGS) $(MAINDIR)

run: build
	@$(TARGET)

build: check
	@go build $(LDFLAGS) -o $(TARGET) $(MAINDIR)

check: clean test

clean:
	@rm -rf bin/

uninstall: clean
	@rm -v $(shell which $(APPNAME))

test: clean
	@go test -v $(shell go list ./... | grep -v /mocks/) -count=1
