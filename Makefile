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
	@echo "build .................................. Start the app container."
	@echo "flush .................................. Remove all instances of the application."
	@echo "status ................................. Display the status of all containers."
	@echo "stop ................................... Destroy the application container."
	@echo " "
	@echo "----------------------------------------------"

build:
	go mod tidy && \
	docker compose up app --build

build\:fresh:
	make flush && \
	docker compose up app --build

flush:
	docker compose down --remove-orphans
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f
	rm -rf ./database/data

stop:
	docker-compose down --volumes

status:
	docker-compose ps
