.PHONY: docker

EXECUTABLE ?= drone-datadog
IMAGE ?= urbint/drone-datadog

docker:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(EXECUTABLE)
	docker build --rm -t $(IMAGE) .
