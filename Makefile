VERSION=$(shell git describe --tags --dirty --always)

LDFLAGS += -extldflags '-static' -w
LDFLAGS += -X github.com/jhoblitt/expfake/version.Version=$(VERSION)

.PHONY: all
all:
	CGO_ENABLED=0 go build -ldflags "${LDFLAGS}"

.PHONY: lint
lint:
	golangci-lint run
