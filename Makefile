.PHONY: help
help: ## Show the help
	@grep -hE '^[A-Za-z0-9_ \-]*?:.*##.*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

GOLANG_VERSION=1.24
GOTENBERG_VERSION=8.21.1
APP_NAME=app
APP_VERSION=snapshot
APP_AUTHOR=app-author
APP_REPOSITORY=https://my.app.git
DOCKER_REGISTRY=app
DOCKER_REPOSITORY=gotenberg

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

.PHONY: fmt
fmt: ## Format Golang codebase and "optimize" the dependencies
	golangci-lint fmt
	go mod tidy
