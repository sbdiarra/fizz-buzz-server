GIN_MODE="release"

default: build

build:
	go build ./...

test:
	go test ./... -v -coverprofile=coverage.out -count=1

run:
	GIN_MODE="release" go run main.go

.PHONY: build test
