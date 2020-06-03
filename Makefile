.PHONY: test
test:
test-all: 
	go test -v ./...

test-int:
	go test -v ./... -run "Int"

test-hello:
	go test -v ./... -run "Hello"

test-gcs:
	go test -v ./pkg/gcs -run 'Bucket'

test-file:
	@CONFIG_FILE=${CURDIR}/configs/config-test.yml go test -v ./pkg/general/... --tags=file_test

test-config:
	@CONFIG_FILE=${CURDIR}/configs/config-test.yml GOCACHE=off  go test -v ./pkg/general/... --tags=config_test
