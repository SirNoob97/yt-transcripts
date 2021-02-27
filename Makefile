MAINDIR := $(shell find . -type f -name 'main.go' -printf '%h\n')
APPNAME := $(shell basename $(MAINDIR))
TARGET := $(shell echo ${PWD}/bin/$(APPNAME))

VERSION := $(shell git describe --tags --abbrev=0)
BUILD := $(shell git rev-list -1 $(VERSION))
LDFLAGS := -ldflags "-X=main.Version=$(VERSION) -X=main.AppName=$(APPNAME)"


.PHONY: install build run clean uninstall

install:
	@echo "Installing the CLI tool"
	@go install $(LDFLAGS) $(MAINDIR)

run: build
	@$(TARGET)

build:
	@echo "Building app binary"
	@go build $(LDFLAGS) -o $(TARGET) $(MAINDIR)

clean:
	@rm -rf bin/

uninstall: clean
	@rm -v $(shell which $(APPNAME))
