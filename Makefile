.PHONY: precommit test-unit
precommit:
	gofmt -w -s -d .
	go vet .
	golangci-lint run
	go mod tidy
	go mod verify
test-unit:
	go test -race -cover ./...
run:
	go run -race main.go
