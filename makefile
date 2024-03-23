# Basic Go command
GOCMD=go

# Program name
PROGRAM_NAME=proxima

# Binary names for each platform and architecture
BINARY_NAME_LINUX_AMD64=$(PROGRAM_NAME)-linux-amd64
BINARY_NAME_LINUX_ARM64=$(PROGRAM_NAME)-linux-arm64
BINARY_NAME_WINDOWS=$(PROGRAM_NAME)-windows-amd64.exe
BINARY_NAME_MAC_AMD64=$(PROGRAM_NAME)-darwin-amd64
BINARY_NAME_MAC_ARM64=$(PROGRAM_NAME)-darwin-arm64

# Target directory for binaries
DIST_DIR=dist

# Builds for all platforms
all: build-linux-amd64 build-linux-arm64 build-windows build-mac-amd64 build-mac-arm64

# Command to create the dist/ directory
$(DIST_DIR):
	mkdir -p $(DIST_DIR)

# Linux AMD64 build
build-linux-amd64: $(DIST_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -o $(DIST_DIR)/$(BINARY_NAME_LINUX_AMD64) main.go

# Linux ARM64 build
build-linux-arm64: $(DIST_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOCMD) build -o $(DIST_DIR)/$(BINARY_NAME_LINUX_ARM64) main.go

# Windows build
build-windows: $(DIST_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOCMD) build -o $(DIST_DIR)/$(BINARY_NAME_WINDOWS) main.go

# macOS AMD64 build
build-mac-amd64: $(DIST_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOCMD) build -o $(DIST_DIR)/$(BINARY_NAME_MAC_AMD64) main.go

# macOS ARM64 build (for M-series chips)
build-mac-arm64: $(DIST_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOCMD) build -o $(DIST_DIR)/$(BINARY_NAME_MAC_ARM64) main.go

# Clean up binaries
clean:
	rm -rf $(DIST_DIR)/*

.PHONY: all build-linux-amd64 build-linux-arm64 build-windows build-mac-amd64 build-mac-arm64 clean

