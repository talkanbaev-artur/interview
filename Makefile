IMAGE_NAME = talkanbaev/test
VERSION = 0.1

build: mod_tidy create_docker tag_latest

run_d: mod_tidy build run_docker
run: 
	go run main.go

create_docker:
	docker build --tag $(IMAGE_NAME):$(VERSION) .

tag_latest:
	docker image tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest

mod_tidy:
	go mod tidy -go=1.16 && go mod tidy -go=1.17

run_docker:
	docker run --name talkanbaev-test --network=host -d $(IMAGE_NAME):$(VERSION)

migrate:
	$(shell chmod +x db/migrate.sh  && bash ./db/migrate.sh migrate)
rollback: 