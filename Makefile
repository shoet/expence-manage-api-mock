.PHONY: help run build
.DEFAULT_GOAL := help

DOCKER_TAG := latest
DOCKER_IMAGE := expense-mock-server
build: ## Build docker image to deploy
	docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} ./

run: ## Do docker container run with hot reload
	docker run --rm -it -p 8080:8080 -v ${PWD}:/app ${DOCKER_IMAGE}:${DOCKER_TAG}

help: ## Show options
	@grep -E '^[a-zA-Z_]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	
