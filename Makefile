.PHONY: default usage start stop

usage:
	@echo "Please provide an option:"
	@echo "	make start		--- Build images and start containers"
	@echo "	make stop		--- Stop running containers"

start:
	DOCKER_BUILDKIT=1 docker build -t tor-loadbalancer .
	DOCKER_BUILDKIT=1 docker build -t tor-circuit tor-circuit/.
	docker-compose up -d

stop:
	docker-compose down

default: usage
