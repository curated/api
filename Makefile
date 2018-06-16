default: test

.PHONY: dist test run

dist:
	@docker build -t curated .

test:
	@go test ./...

run:
	@go run main.go -logtostderr=true
