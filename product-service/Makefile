# ----------------------------
# Env Variables
# ----------------------------
CONTAINER_NAME ?= "product-service"
DATABASE_CONTAINER ?= database
API_CONTAINER ?= server
DOCKER_COMPOSE_FILE ?= docker-compose.yaml

## run: starts containers to run api server
run: api-create

## setup: executes pre-defined steps to setup api server
setup: pg-create

## api-create: starts api server
api-create:
	@echo Starting Api container
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p=${CONTAINER_NAME} up ${API_CONTAINER}
	@echo Api container started!

## pg-create: starts postgres container
pg-create:
	@echo Starting Postgres database container
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d ${DATABASE_CONTAINER}
	@echo Database container started!

## pg-drop: reset db to blank
pg-drop:
	@echo Dropping database...
	docker-compose --profile tools run --rm pg-migrate drop
	@echo Done!

## down: removes all docker containers
down:
	@echo Removing containers...
	docker-compose down
	@echo Done!
