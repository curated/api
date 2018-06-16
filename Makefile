default: test

.PHONY: dist test run

dist:
	@docker build -t curated .

test:
	@CONFIG=config/config.test.json go test ./...

run:
	@go run main.go -logtostderr=true
