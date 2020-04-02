.PHONY: precommit test-unit gen-proto run test test-integration up down restart

DOCKER_COMPOSE_FILE ?= deployments/docker-compose/docker-compose.yml
DOCKER_COMPOSE_TEST_FILE ?= deployments/docker-compose/docker-compose.test.yml

precommit:
	gofmt -w -s -d .
	go vet .
	golangci-lint run --timeout=30m
	go mod tidy
	go mod verify

test-unit:
	go test -race -cover ./internal/bucket/...

test-integration:
	docker-compose -f ${DOCKER_COMPOSE_TEST_FILE} up --build -d ;\
		docker-compose -f ${DOCKER_COMPOSE_TEST_FILE} run integration_tests go test ./internal/integration-tests;\
		test_status_code=$$? ;\
		docker-compose -f ${DOCKER_COMPOSE_TEST_FILE} down ;\
		exit $$test_status_code ;\

gen-proto:
	 protoc -I. api/antibruteforce.proto --go_out=plugins=grpc:internal/antibruteforce/delivery/grpc

run:
	go run -race main.go serve

up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} up

down:
	docker-compose -f ${DOCKER_COMPOSE_FILE} down

restart: down up

test: test-unit test-integration

