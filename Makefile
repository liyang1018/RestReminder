# Makefile for building Timer application for macOS

# Compiler and compiler flags
GO=go
FYNE=/Users/vito.li/GolandProjects/gopath/bin/fyne
APP_NAME=Timer
OS=darwin

# Main build target
build:
	@echo "Building $(APP_NAME) for $(OS)"
	$(FYNE) package -os $(OS) -name $(APP_NAME)

# Clean target
clean:
	@echo "Cleaning up"
	rm -rf $(APP_NAME)

# Help target
help:
	@echo "Usage: make [build|clean|help]"
	@echo "  build:     Build $(APP_NAME) for $(OS)"
	@echo "  clean:     Remove build artifacts"
	@echo "  help:      Show this help message"

# Default target
.DEFAULT_GOAL := help
