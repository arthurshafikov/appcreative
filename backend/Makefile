DOCKER_COMPOSE_FILE := "./deployments/docker-compose.yml"
APP_NAME := "appcreative"

test: 
	go test --short -race ./internal/...

.PHONY: test

up:
	docker-compose --env-file ./.env -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} up -d --build

down:
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} down --volumes
