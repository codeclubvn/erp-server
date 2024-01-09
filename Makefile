CONTAINER_NAME=erp_server
CONTAINER_PORT=8008
IS_EXIST := $(shell docker ps -a --filter "name=$(CONTAINER_NAME)" --format '{{.Names}}')
NETWORK=oplacrm

# Build Docker image and run container
build:
	docker compose down
	docker rmi $(CONTAINER_NAME)
	docker compose up -d

# Run the Golang server
run:
	go run main.go

pull:
	git pull origin develop
