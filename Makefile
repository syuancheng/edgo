.PHONY: fmt vet test race check

GO_CACHE ?= /tmp/edgo-gocache

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

vet:
	GOCACHE=$(GO_CACHE) go vet ./...

test:
	GOCACHE=$(GO_CACHE) go test ./...

race:
	GOCACHE=$(GO_CACHE) go test -race ./...

check:
	@test -z "$$(gofmt -l $$(find . -name '*.go' -not -path './vendor/*'))" || \
		(echo 'Go files need formatting; run make fmt' && exit 1)
	$(MAKE) vet
	$(MAKE) test
