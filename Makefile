IMAGE_NAME = talkanbaev/test
VERSION = 0.1

build: mod_tidy create_docker tag_latest

run: build run_docker

create_docker:
	docker build --tag $(IMAGE_NAME):$(VERSION) .

tag_latest:
	docker image tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest

mod_tidy:
	go mod tidy

run_docker:
	docker run --name talkanbaev-test -d $(IMAGE_NAME):$(VERSION)
