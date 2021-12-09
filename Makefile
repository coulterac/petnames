.DEFAULT_GOAL := all
all: security test vendor vet

security:
	go get github.com/securego/gosec/cmd/gosec/...
	gosec ./...

test:
	go test -v

vendor:
	go mod vendor

vet:
	go vet ./...

.PHONY: all test security vendor vet