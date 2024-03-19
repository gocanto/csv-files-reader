.PHONY: gocanto\:trades

include .env

DB_NETWORK = gocanto
APP_PATH = $(shell pwd)

help:
	@echo " "
	@echo "----------------------------------------------"
	@echo "               Api CLI Help                   "
	@echo "----------------------------------------------"
	@echo " "
	@echo "run:api ................................ Start the application."
	@echo "build .................................. Build the app container."
	@echo "build:fresh ............................ Build a fresh instance of the app container."
	@echo "status ................................. Display the status of all containers."
	@echo "stop ................................... Destroy the application container."
	@echo " "
	@echo "----------------------------------------------"


run\:api:
	cd cmd/app && go mod tidy && go mod download && \
	CGO_ENABLED=0 go run -tags app github.com/ohlc-price-data/cmd/app

build:
	docker compose up app --build

build\:fresh:
	docker compose down --remove-orphans && \
	docker container prune -f && \
	docker image prune -f && \
	docker volume prune -f && \
	docker network prune -f && \
	docker compose up app --build

stop:
	docker-compose down --volumes

status:
	docker-compose ps
