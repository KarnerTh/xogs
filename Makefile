GIT_TAG := $(shell git describe --tags --abbrev=0)

.PHONY: run
run:
	go run main.go

# https://github.com/mfridman/tparse is needed for pretty output
.PHONY: test
test:
	go test ./... -json | tparse -all

.PHONY: test-cleanup
test-cleanup:
	go clean -testcache

.PHONY: publish
publish:
	GOPROXY=proxy.golang.org go list -m github.com/KarnerTh/xogs@$(GIT_TAG)
