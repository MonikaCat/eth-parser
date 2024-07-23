###############################################################################
###                                  All                                    ###
###############################################################################
all: lint build install

###############################################################################
###                                 Build                                   ###
###############################################################################
build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "Building eth-parser for Windows..."
	@go build -mod=readonly -o build/eth-parser.exe ./cmd/eth-parser
else
	@echo "Building eth-parser for Unix-like OS..."
	@go build -mod=readonly -o build/eth-parser ./cmd/eth-parser
endif
.PHONY: build

###############################################################################
###                                Install                                  ###
###############################################################################
install: go.sum
	@echo "Installing eth-parser..."
	@go install -mod=readonly ./cmd/eth-parser
.PHONY: install

###############################################################################
###                                 Tests                                 ###
###############################################################################
clean:
	@echo "Cleaning build..."
	@rm -rf ./build/**
.PHONY: clean

test-unit: 
	@echo "Executing unit tests..."
	@go test -mod=readonly -v -coverprofile=coverage.txt ./...
	@echo "Unit tests completed."
	@exit 0  # Ensure the script exits cleanly
.PHONY: test-unit


###############################################################################
###                                 Lint                                    ###
###############################################################################
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	@echo "Running linter..."
	@go run $(golangci_lint_cmd) run --timeout=10m

lint-fix:
	@echo "Running linter..."
	@go run $(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0

.PHONY: lint lint-fix

.PHONY: lint lint-fix format
###############################################################################
###                                 Help                                  ###
###############################################################################
help:
	@echo "Available commands:"
	@echo "  make all        	: Run lint and build"
	@echo "  make build      	: Build the eth-parser binary"
	@echo "  make install    	: Install the eth-parser binary"
	@echo "  make clean      	: Clean the eth-parser build"
	@echo "  make lint       	: Run linter"
	@echo "  make lint-fix   	: Run linter with auto-fixes"
	@echo "  make help       	: Print this help message"
.PHONY: help