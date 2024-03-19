.PHONY: gocanto\:trades

include .env

DB_NETWORK = gocanto
APP_PATH = $(shell pwd)

run\:api:
	cd cmd/app && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run -tags app github.com/ohlc-price-data/cmd/app

build:
	docker compose up app --build

build\:fresh:
	make docker\:prune && \
	docker compose up app --build

docker\:prune:
	docker compose down --remove-orphans
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f
