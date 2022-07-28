BUILD_TARGET = wolper
IMAGE_NAME ?= ghcr.io/peng225/wolper

GO_FILES:=$(shell find . -type f -name '*.go' -print)

$(BUILD_TARGET): $(GO_FILES)
	CGO_ENABLED=0 go build -o $@ -v

.PHONY: proto
proto:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/wolper.proto

.PHONY: test
test: $(BUILD_TARGET)
	go test -v `go list ./... | grep -v helm_test`

.PHONY: image
image:
	docker build . --file Dockerfile --tag $(IMAGE_NAME)

.PHONY: clean
clean:
	rm -f $(BUILD_TARGET)
