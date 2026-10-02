.PHONY: start form stat base

start:
	@go run ./cmd/api

form:
	@gofmt -w .

stat:
	@git status

base:
	@git add . && git commit -m "$(msg)" && git push


dock:
	@docker run -d --rm --name logsence -p 8080:8080 logsence

in:
	@docker exec -it d1dd4fae7316 sh
