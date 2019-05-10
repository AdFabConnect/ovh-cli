PROJECT_NAME := ovh-cli
PKG := github.com/AdFabConnect/ovh-cli
OUTPUT_DIR := output

GO_VERSION := $(shell go version | awk '{print $$3}' )
APP_VERSION := $(shell git describe --tag >/dev/null 2>&1; if [ $$? -ne 0 ]; then git rev-parse --short HEAD; else git describe --tag; fi)
GIT_COMMIT := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell date '+%Y-%m-%d_%H:%M:%S' )

LDFLAGS = '-X ${PKG}/cmd.Version=${APP_VERSION} -X ${PKG}/cmd.GoVersion=${GO_VERSION} -X ${PKG}/cmd.OsArchi=${GOOS}/${GOARCH} -X ${PKG}/cmd.GitCommit=${GIT_COMMIT} -X ${PKG}/cmd.BuildDate=${BUILD_DATE}'
OUTPUT = ${OUTPUT_DIR}/${GOOS}-${GOARCH}/${PROJECT_NAME}

.PHONY: all
all: build

.PHONY: build
build:
	$(eval GOOS := $(shell go env GOOS))
	$(eval GOARCH := $(shell go env GOARCH))
	@echo "Build and install ${PROJECT_NAME} - ${GOOS} ${GOARCH}"
	@go build -o ${OUTPUT} ${FLAGS} -ldflags ${LDFLAGS}


.PHONY: clean
clean:
	rm -rf ${OUTPUT_DIR}
