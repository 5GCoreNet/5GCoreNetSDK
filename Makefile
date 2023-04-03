GOPATH := $(shell go env GOPATH)

dependencies:
	echo "Installing dependencies"
	go install github.com/golang/mock/mockgen@v1.6.0
	go mod download

mock-gen: dependencies
	echo "Generating mock files"
	PATH=$(PATH):$(GOPATH)/bin go generate ./...

test: mock-gen
	echo "Running tests"
	go test -v ./...

.PHONY: mock-gen
