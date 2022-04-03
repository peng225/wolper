BUILD_TARGET = wolper
IMAGE_NAME ?= wolper
INPUT ?= input
TAG ?= v0.3.1

$(BUILD_TARGET):
	go build -o $@ -v

.PHONY: test
test: $(BUILD_TARGET)
	go test -v ./...

.PHONY: image
image:
	docker build . --file Dockerfile --tag $(IMAGE_NAME):$(TAG) --build-arg INPUT=$(INPUT) --build-arg TAG=$(TAG)

.PHONY: clean
clean:
	rm -f $(BUILD_TARGET)
	docker rmi $(IMAGE_NAME):$(TAG)
