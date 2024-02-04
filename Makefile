GO ?= go
GOLANGCILINT ?= golangci-lint

BINARY := version
REPOSITORY ?= ghcr.io/kvanzuijlen/version

GO_MAJOR_VERSION = $(shell $(GO) version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell $(GO) version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
MINIMUM_SUPPORTED_GO_MAJOR_VERSION = 1
MINIMUM_SUPPORTED_GO_MINOR_VERSION = 21

.PHONY: lint
lint: validate-go-version
	GO111MODULE=on $(GOLANGCILINT) run

.PHONY: clean
clean:
	-rm -rf release
	-rm -f $(BINARY)

.PHONY: build
build: validate-go-version clean $(BINARY)

$(BINARY):
	CGO_ENABLED=0 $(GO) build -a -installsuffix cgo -ldflags="-X github.com/kvanzuijlen/version/cmd.VERSION=${VERSION}" -o $@

DOCKER_BUILD_PLATFORM         ?= linux/amd64,linux/arm64,linux/ppc64le,linux/arm/v7
DOCKER_BUILDX_ARGS_LIST       ?= \
	CREATED=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ') \
	REVISION=$(shell git rev-parse --short HEAD) \
	VERSION=${VERSION}
DOCKER_BUILDX_ARGS            := $(addprefix --build-arg ,$(DOCKER_BUILDX_ARGS_LIST))
DOCKER_BUILDX_TAGS			  := -t $(REPOSITORY):latest -t $(REPOSITORY):${VERSION}
DOCKER_BUILDX                 := docker buildx build ${DOCKER_BUILDX_ARGS} --platform ${DOCKER_BUILD_PLATFORM} ${DOCKER_BUILDX_TAGS}
DOCKER_BUILDX_PUSH            := $(DOCKER_BUILDX) --push

.PHONY: docker-build
docker-build:
	$(DOCKER_BUILDX) .

.PHONY: docker-push
docker-push:
	$(DOCKER_BUILDX_PUSH) .

.PHONY: validate-go-version
validate-go-version:
	@if [ $(GO_MAJOR_VERSION) -gt $(MINIMUM_SUPPORTED_GO_MAJOR_VERSION) ]; then \
		exit 0 ;\
	elif [ $(GO_MAJOR_VERSION) -lt $(MINIMUM_SUPPORTED_GO_MAJOR_VERSION) ]; then \
		echo '$(GO_VERSION_VALIDATION_ERR_MSG)';\
		exit 1; \
	elif [ $(GO_MINOR_VERSION) -lt $(MINIMUM_SUPPORTED_GO_MINOR_VERSION) ] ; then \
		echo '$(GO_VERSION_VALIDATION_ERR_MSG)';\
		exit 1; \
	fi
