SHELL = /bin/sh

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=api

all: test build
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/finance-api-server/main.go
test: 
		$(GOTEST) -v ./...
clean: 
		rm -f ./$(BINARY_NAME)
		rm -f ./$(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/finance-api-server/main.go
		./$(BINARY_NAME) --scheme=http --port 3000
