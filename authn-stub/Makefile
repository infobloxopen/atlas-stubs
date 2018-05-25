PROJECT_ROOT		:= github.com/infobloxopen/atlas-stubs/authn-stub
BUILD_PATH  		:= bin
DOCKERFILE_PATH		:= $(CURDIR)/docker

# configuration for image names
USERNAME        := $(USER)
GIT_COMMIT      := $(shell git describe --dirty=-unsupported --always || echo pre-commit)
IMAGE_VERSION   ?= $(USERNAME)-dev-$(GIT_COMMIT)
IMAGE_REGISTRY  ?= infoblox

# configuration for gateway binary and image
BINARY 	   := $(BUILD_PATH)/authn-stub
SRC_PATH := $(PROJECT_ROOT)/cmd/authn-stub
IMAGE		   := $(IMAGE_REGISTRY)/authn-stub
DOCKERFILE := $(DOCKERFILE_PATH)/Dockerfile

# configuration for building on host machine
GO_PACKAGES		:= $(shell go list ./... | grep -v vendor)

.PHONY: all
all: vendor authn-stub-docker

.PHONY: fmt
fmt:
	@go fmt $(GO_PACKAGES)

.PHONY: authn-stub-docker
authn-stub-docker:
	@docker build -f $(DOCKERFILE) -t $(IMAGE):$(IMAGE_VERSION) .

.PHONY: push
push:
	@docker push $(IMAGE)

.PHONY: vendor
vendor:
	@dep ensure -vendor-only

.PHONY: vendor-update
vendor-update:
	@dep ensure

.PHONY: clean
clean:
	@docker rmi -f $(shell docker images -q $(IMAGE)) || true
	@docker rmi -f $(shell docker images --filter "label=intermediate=true" -q) || true
