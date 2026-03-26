NAME := goscan
VERSION := 3.0
AUTHOR := Marco Lancini [@LanciniMarco] - Community Contributors
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.author=$(AUTHOR)'

.DEFAULT_GOAL := help

.PHONY: init
init:  ## Init the project.
	go get -u github.com/golang/dep/cmd/dep
	dep init

.PHONY: setup
setup:  ## Setup for required tools.
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

.PHONY: fmt
fmt: ## Formatting source codes.
	go fmt ./...

.PHONY: lint
lint: ## Run golint and go vet.
	@golint ./core/...
	@go vet ./core/...

.PHONY: build
build: main.go  ## Build for current system.
	go build -ldflags "$(LDFLAGS)" -o goscan
	@echo "Build successful: ./goscan"

.PHONY: build-windows
build-windows: main.go  ## Build for Windows (AMD64).
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o goscan.exe
	@echo "Windows binary built: ./goscan.exe"

.PHONY: build-macos
build-macos: main.go  ## Build for macOS (AMD64).
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o goscan_macos
	@echo "macOS binary built: ./goscan_macos"

.PHONY: build-kali
build-kali: main.go  ## Build for Kali Linux (AMD64).
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o goscan_kali
	@echo "Kali Linux binary built: ./goscan_kali"

.PHONY: build-rpi
build-rpi: main.go  ## Build for Raspberry Pi 5 (ARM64).
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o goscan_rpi5
	@echo "Raspberry Pi 5 binary built: ./goscan_rpi5"

.PHONY: build-rpi-32
build-rpi-32: main.go  ## Build for older Raspberry Pi (ARMv7).
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "$(LDFLAGS)" -o goscan_rpi_armv7
	@echo "Raspberry Pi ARMv7 binary built: ./goscan_rpi_armv7"

.PHONY: cross
cross: main.go  ## Build for all platforms.
	mkdir -p pkg
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o goscan.exe
	zip pkg/goscan_$(VERSION)_windows_amd64.zip goscan.exe
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o goscan
	zip pkg/goscan_$(VERSION)_darwin_amd64.zip goscan
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o goscan
	zip pkg/goscan_$(VERSION)_linux_amd64.zip goscan
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o goscan
	zip pkg/goscan_$(VERSION)_rpi5_arm64.zip goscan
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "$(LDFLAGS)" -o goscan
	zip pkg/goscan_$(VERSION)_rpi_armv7.zip goscan
	@echo "All cross-platform binaries built in pkg/"

.PHONY: install
install: build  ## Install goscan (Linux/macOS only).
	sudo cp goscan /usr/local/bin/goscan
	sudo chmod +x /usr/local/bin/goscan
	@echo "GoScan installed to /usr/local/bin/goscan"

.PHONY: uninstall
uninstall:  ## Uninstall goscan (Linux/macOS only).
	sudo rm -f /usr/local/bin/goscan
	@echo "GoScan uninstalled"

.PHONY: run
run: build  ## Build and run goscan.
	sudo ./goscan

.PHONY: docker-build
docker-build:  ## Build Docker image.
	docker build -f ../cli/Dockerfile -t goscan:latest .

.PHONY: docker-run
docker-run:  ## Run GoScan in Docker.
	docker run -it goscan:latest

.PHONY: clean
clean:  ## Clean build artifacts.
	rm -f goscan goscan.exe goscan_*
	rm -rf pkg/
	@echo "Clean complete"

.PHONY: help
help: ## Show this help message.
	@echo "GoScan v$(VERSION) - Advanced Network Scanner"
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    %-20s %s\n", $$1, $$2}'
