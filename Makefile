# Makefile

# The name of the executable (default is current directory name)
TARGET := filscan-delay-monitor

# Go related variables.
GOBASE := $(shell pwd)

# Build the project
all: clean build

build:
	@echo "  >  Building binary..."
	go build -mod=mod -o $(GOBASE)/$(TARGET) $(GOBASE)/$(TARGET).go

clean:
	@echo "  >  Cleaning build cache"
	go clean -mod=mod

.PHONY: all build clean