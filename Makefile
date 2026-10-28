.PHONY: help
help: ## Show the help
	@grep -hE '^[A-Za-z0-9_ \-]*?:.*##.*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

GOLANG_VERSION=1.27.3
GOTENBERG_VERSION=8.39.0
APP_NAME=app
APP_VERSION=snapshot
APP_AUTHOR=app-author
APP_REPOSITORY=https://my.app.git
DOCKER_REGISTRY=app
DOCKER_REPOSITORY=gotenberg

# golangci-lint runs at a pinned version, so a Homebrew upgrade cannot move the
# findings under the zero-warnings policy. An env var of the same name overrides
# it.
GOLANGCI_LINT_VERSION ?= v2.15.0
GOLANGCI_LINT = go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

.PHONY: build
build: ## Build the Gotenberg's Docker image
	docker build \
	--build-arg GOLANG_VERSION=$(GOLANG_VERSION) \
	--build-arg GOTENBERG_VERSION=$(GOTENBERG_VERSION) \
	--build-arg APP_NAME=$(APP_NAME) \
	--build-arg APP_VERSION=$(APP_VERSION) \
	--build-arg APP_AUTHOR=$(APP_AUTHOR) \
	--build-arg APP_REPOSITORY=$(APP_REPOSITORY) \
	-t $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):$(GOTENBERG_VERSION)-$(APP_NAME)-$(APP_VERSION) \
	-f build/Dockerfile .

.PHONY: lint
lint: ## Lint Golang codebase
	$(GOLANGCI_LINT) run

.PHONY: lint-prettier
lint-prettier: ## Lint non-Golang codebase
	npx prettier --check .

.PHONY: lint-todo
lint-todo: ## Find TODOs in Golang codebase
	$(GOLANGCI_LINT) run --no-config --default=none --enable godox

.PHONY: fmt
fmt: ## Format Golang codebase and "optimize" the dependencies
	go fix ./...
	$(GOLANGCI_LINT) fmt
	go mod tidy

.PHONY: prettify
prettify: ## Format non-Golang codebase
	npx prettier --write .
