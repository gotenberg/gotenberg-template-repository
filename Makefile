.PHONY: help
help: ## Show the help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: it
it: build build-tests ## Initialize the development environment

GOLANG_VERSION=1.24
GOTENBERG_VERSION=8.18.0
APP_NAME=app
APP_VERSION=snapshot
APP_AUTHOR=app-author
APP_REPOSITORY=https://my.app.git
DOCKER_REGISTRY=app
DOCKER_REPOSITORY=gotenberg
GOLANGCI_LINT_VERSION=v1.64.2 # See https://github.com/golangci/golangci-lint/releases.

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

.PHONY: build-tests
build-tests: ## Build the tests' Docker image
	docker build \
	--build-arg GOLANG_VERSION=$(GOLANG_VERSION) \
	--build-arg DOCKER_REGISTRY=$(DOCKER_REGISTRY) \
	--build-arg DOCKER_REPOSITORY=$(DOCKER_REPOSITORY) \
	--build-arg GOTENBERG_VERSION=$(GOTENBERG_VERSION) \
	--build-arg APP_NAME=$(APP_NAME) \
	--build-arg APP_VERSION=$(APP_VERSION) \
	--build-arg GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION) \
	-t $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):$(GOTENBERG_VERSION)-$(APP_NAME)-$(APP_VERSION)-tests \
	-f test/Dockerfile .

.PHONY: tests
tests: ## Start the testing environment
	docker run --rm -it \
	-v $(PWD):/tests \
	$(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):$(GOTENBERG_VERSION)-$(APP_NAME)-$(APP_VERSION)-tests \
	bash

.PHONY: tests-once
tests-once: ## Run the tests once (prefer the "tests" command while developing)
	docker run --rm  \
	-v $(PWD):/tests \
	$(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):$(GOTENBERG_VERSION)-$(APP_NAME)-$(APP_VERSION)-tests \
	gotest

# go install mvdan.cc/gofumpt@latest
# go install github.com/daixiang0/gci@latest
.PHONY: fmt
fmt: ## Format the code and "optimize" the dependencies
	gofumpt -l -w .
	# TODO: change github.com/gotenberg/gotenberg-template-repository to actual namespace.
	gci write -s standard -s default -s "prefix(github.com/gotenberg/gotenberg/v8)" -s "prefix(github.com/gotenberg/gotenberg-template-repository)" --skip-generated --skip-vendor --custom-order .
	go mod tidy
