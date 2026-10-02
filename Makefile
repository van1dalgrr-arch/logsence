APP := logsence

.PHONY: help start form stat base \
        build dock in logs stop \
        up down ps clogs db reset

help:
	@echo "Go:"
	@echo "  start   - запустить локально (go run)"
	@echo "  form    - отформатировать код (gofmt)"
	@echo ""
	@echo "Git:"
	@echo "  stat    - git status"
	@echo "  base    - add + commit + push  (make base msg=\"...\")"
	@echo ""
	@echo "Docker:"
	@echo "  build   - собрать образ"
	@echo "  dock    - запустить контейнер в фоне"
	@echo "  in      - зайти внутрь контейнера"
	@echo "  logs    - логи контейнера"
	@echo "  stop    - остановить контейнер"
	@echo ""
	@echo "Compose:"
	@echo "  up      - поднять app + postgres"
	@echo "  down    - остановить всё"
	@echo "  ps      - статус сервисов"
	@echo "  clogs   - логи приложения"
	@echo "  db      - консоль postgres (psql)"
	@echo "  reset   - удалить всё вместе с данными базы"

# go
start:
	@go run ./cmd/api

form:
	@gofmt -w .

# git
stat:
	@git status

base:
	@git add . && git commit -m "$(msg)" && git push

# Docker
build:
	@docker build -t $(APP) .

dock:
	@docker run -d --rm --name $(APP) -p 8080:8080 $(APP)

in:
	@docker exec -it $(APP) sh

logs:
	@docker logs -f $(APP)

stop:
	@docker stop $(APP)

# Compose
up:
	@docker compose up -d --build

down:
	@docker compose down

ps:
	@docker compose ps

clogs:
	@docker compose logs -f app

db:
	@docker compose exec postgres psql -U postgres -d $(APP)

reset:
	@docker compose down -v
