.PHONY: test
test:
test-all: 
	go test -v ./...

test-int:
	go test -v ./... -run "Int"

test-hello:
	go test -v ./... -run "Hello"