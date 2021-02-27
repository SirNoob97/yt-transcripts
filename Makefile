.PHONY: client

client:
	@echo "Building binary"
	go build -o bin/client cmd/client/main.go
