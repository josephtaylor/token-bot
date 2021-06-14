REGISTRY_IMAGE := josephtaylor/token-bot
REGISTRY_REPO := github.com
PACKAGE := token-bot
SCM_URL :=

.PHONY: help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := artifact

build: clean
	docker buildx build --build-arg VERSION=$(NEXT_VERSION),PACKAGE=$(PACKAGE) -t $(REGISTRY_IMAGE):latest --load .

artifact: build
	mkdir -p .target/bin
	docker create --name token-extraction $(REGISTRY_IMAGE):latest
	docker cp token-extraction:/token .target/bin
	docker rm token-extraction

push: build
	docker tag $(REGISTRY_IMAGE):latest $(IMAGE_PATH_TAGGED_SHA)
	docker push $(IMAGE_PATH_TAGGED_SHA)

tag:
	@eval git remote set-url origin $(SCM_URL)
	@eval ./version -f patch

release:
	docker pull $(IMAGE_PATH_TAGGED_SHA)
	docker tag $(IMAGE_PATH_TAGGED_SHA) $(IMAGE_PATH_UNTAGGED):$(CURR_VERSION)
	docker tag $(IMAGE_PATH_TAGGED_SHA) $(IMAGE_PATH_UNTAGGED):latest
	docker push $(IMAGE_PATH_UNTAGGED):$(CURR_VERSION)
	docker push $(IMAGE_PATH_UNTAGGED):latest

gotools:
	go get golang.org/x/lint/golint

gofmt:
	go fmt ./...

golint:
	golint ./...

vet:
	go vet -v ./...

clean-dist:
	rm -Rf dist/*

clean: clean-dist
	go clean -i ./...
	docker rm token-extraction | true

clean-vendor:
	find $(CURDIR)/vendor -type d -print0 2>/dev/null | xargs -0 rm -Rf

govendor:
	go mod vendor