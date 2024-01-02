CONTAINER_NAME=erp_server
CONTAINER_PORT=8008
IS_EXIST := $(shell docker ps -a --filter "name=$(CONTAINER_NAME)" --format '{{.Names}}')
NETWORK=oplacrm

# Build Docker image and run container
build:
	@if "$(CONTAINER_NAME)" == "$(IS_EXIST)" ( docker rm -f $(CONTAINER_NAME) && docker rmi $(CONTAINER_NAME) ) else ( echo "Container $(CONTAINER_NAME) not exist" )
	docker build -t $(CONTAINER_NAME) .
	docker run -d -p $(CONTAINER_PORT):$(CONTAINER_PORT) --name $(CONTAINER_NAME) --network $(NETWORK) $(CONTAINER_NAME)

# Run the Golang server
run:
	go run main.go

pull:
	git pull origin develop
