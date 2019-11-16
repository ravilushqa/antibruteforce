.PHONY: precommit test-unit gen-proto run
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

