.PHONY: precommit test-unit gen-proto run

DOCKER_COMPOSE_FILE ?= deployments/docker-compose/docker-compose.yml
DOCKER_COMPOSE_TEST_FILE ?= deployments/docker-compose/docker-compose.test.yml

precommit:
	gofmt -w -s -d .
	go vet .
	golangci-lint run
	go mod tidy
	go mod verify

test-unit:
	go test -race -cover ./...

gen-proto:
	 protoc -I. api/antibruteforce.proto --go_out=plugins=grpc:internal/antibruteforce/delivery/grpc

run:
	go run -race main.go serve

up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} up

down:
	docker-compose -f ${DOCKER_COMPOSE_FILE} down

