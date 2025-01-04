.PHONY: build install clean server client b i c

BIN_DIR := ./bin

build:
	go mod tidy
	go build -o $(BIN_DIR)/chat-server ./server/ 
	go build -o $(BIN_DIR)/chat-client ./client/

install: build
	go install ./client
	go install ./server

# Clean up build artifacts
clean:
	@rm -rf $(BIN_DIR)/*

server:
	$(BIN_DIR)/chat-server

client:
	$(BIN_DIR)/chat-client

# Shortcuts
b: build
i: install
c: clean

