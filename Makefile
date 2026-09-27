.PHONY: run test build up down logs stat base form db

start:
	@go run ./cmd/api

form:
	@gofmt -w .

stat:
	git add . && git commit -m "$(msg)" && git push

base:
	@∏git add . && git commit -m $msg
