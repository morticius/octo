BIN_PREFIX?=octo
BIN_DIR?=bin
PORT?=8000
PLATFORMS := linux/amd64 darwin/amd64 windows/amd64
GOPATH_DIR=`go env GOPATH`

.PHONY: clean
clean: ## Clean binary directory
	@rm -f ./${BIN_DIR}/*

.PHONY: tidy
tidy: ## Run go mod tidy
	@go mod tidy -v

.PHONY: build
build: | clean tidy ## Build binaries
	@COMMIT="git-$$(git rev-parse --short HEAD)" ; \
	REPO_INFO=$$(git config --get remote.origin.url) ; \
	VERSION=$$(cat VERSION) ; \
	BUILD_TIME=$$(date +"%Y-%m-%d_%H:%M:%S") ; \
	FULL_PATH="$$(pwd)" ; \
	PKG_PATH="$${FULL_PATH##*src/}" ; \
	for platform in $(PLATFORMS); do \
        GOOS="$${platform%%/*}" ; \
        GOARCH="$${platform##*/}" ; \
        output_name="${BIN_PREFIX}-$${GOOS}-$${GOARCH}" ; \
        LDFLAGS="-s -w -X main.BuildTime=$${BUILD_TIME} -X main.Version=$${VERSION} -X main.Commit=$${COMMIT} -X main.Repo=$${REPO_INFO}" ; \
        env GOOS=$${GOOS} GOARCH=$${GOARCH} go build -v -o bin/$${output_name} -ldflags "$$LDFLAGS" ./cmd/server ; \
        if [[ $$? -ne 0 ]]; then \
            echo "!!! an error has occurred! Aborting the script execution... !!!" ; \
            exit 1 ; \
        else \
            echo "# build bin/$${output_name} done." ; \
        fi \
    done

.PHONY: run-darwin
run-darwin: build ## Run darwin/amd64 binary
	./${BIN_DIR}/${BIN_PREFIX}-darwin-amd64 -c ./configs/server/local.toml

.PHONY: run-linux
run-linux: build ## Run linux/amd64 binary
	./${BIN_DIR}/${BIN_PREFIX}-darwin-amd64 -c ./configs/server/local.toml

.PHONY: run-windows
run-windows: build ## Run windows/amd64 binary
	./${BIN_DIR}/${BIN_PREFIX}-windows-amd64 -c ./configs/server/local.toml

.PHONY: test
test: ## Run tests
	@go test -v -race ./... ; \
	@echo ""

.PHONY: outdated-deps
outdated-deps: ## List outdated dependencies
	@go list -u -m -f '{{if not .Indirect}}{{if .Update}}{{.}}{{end}}{{end}}' all

.PHONY: usage
usage: ## List available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: lint
lint: ## Run linter
	golint ./...

.PHONY: vet
vet: ## Run go vet
	go vet -v ./...

.PHONY: fmt
fmt: ## Run formatting tool for go
	go fmt ./...

.PHONY: env
env: ## Show environment variables
	@go env

.PHONY: install-golangci-lint
install-golangci-lint: ## Install golangci-linter
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH_DIR)/bin v1.41.1

.PHONY: golangci-lint
golangci-lint: ## Install golangci-linter
	@$(GOPATH_DIR)/bin/golangci-lint run

.DEFAULT_GOAL := usage
