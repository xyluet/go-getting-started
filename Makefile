GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/go-getting-started
IMAGE := golang-docker-heroku
VERSION := $(shell git describe --tags --always --dirty)

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web

.PHONY: docker
docker:
	docker build \
		-t $(IMAGE):latest \
		-t $(IMAGE):$(VERSION) \
		-f Dockerfile .

.PHONY: push-heroku
push-heroku:
	docker tag $(IMAGE):$(VERSION) registry.heroku.com/$(IMAGE)/app
	docker push registry.heroku.com/$(IMAGE)/app

.PHONY: release-heroku
release-heroku:
	heroku container:release app
