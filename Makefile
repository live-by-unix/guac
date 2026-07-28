.PHONY: build test clean install lint

BINARY_NAME=guac
BUILD_DIR=build
GO_FILES=$(shell find . -name '*.go' -not -path './vendor/*')

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(BINARY_NAME)

install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

lint:
	@echo "Running linter..."
	@golangci-lint run

run: build
	@$(BUILD_DIR)/$(BINARY_NAME)

deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
