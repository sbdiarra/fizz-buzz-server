IMAGE_TAG=fizzbuzz:1.0
default: build

build:
	go build ./...

build-image:
	docker build -t fizzbuzz:1.0 . 

test:
	go test -mod vendor ./... -v -coverprofile=coverage.out -count=1

run-locally:
	GIN_MODE="release" go run main.go

run-container:
	docker run --name=fizzbuzz -p 8080:8080 $(IMAGE_TAG)


.PHONY: build test
