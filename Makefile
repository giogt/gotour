.PHONY: all
all: build test

.PHONY: build
build:
	go build ./pkg/... ./cmd/...

.PHONY: test
test:
	go test ./pkg/... ./cmd/...

.PHONY: clean
clean:
	go clean
	go clean -testcache
