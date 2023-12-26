help: ## Prints available commands
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[.a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

docker-dev: docker-down ## Start in development mode
	docker-compose -f docker-compose.dev.yml up

docker-dev--build: docker-down ## Start in development mode
	docker-compose up -f docker-compose.dev.yml --build

docker-down: ## Remove containers and volumes
	docker-compose down -v --remove-orphans

docker-down-rmi: ## Removes containers, images and volumes
	docker-compose down -v --remove-orphans --rmi all

docker: ## Start docker compose
	docker-compose up

docker--build: ## Rebuild images and start
	docker-compose up --build
