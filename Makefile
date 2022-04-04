BUILD_TARGET = wolper
IMAGE_NAME ?= wolper

$(BUILD_TARGET):
	CGO_ENABLED=0 go build -o $@ -v

.PHONY: test
test: $(BUILD_TARGET)
	go test -v ./...

.PHONY: image
image:
	docker build . --file Dockerfile --tag $(IMAGE_NAME)

.PHONY: clean
clean:
	rm -f $(BUILD_TARGET)
	docker rmi $(IMAGE_NAME)
