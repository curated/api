default: test

.PHONY: test build run push

test:
	@CONFIG=config/test.config.json go test ./...

build:
	@docker build -t curated .

run:
	@go run main.go -logtostderr=true

push:
	@now
