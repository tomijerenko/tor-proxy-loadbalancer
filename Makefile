.PHONY: default usage build run build-release docker-start

usage:
	@echo "Please provide an option:"
	@echo "	make build		--- Build the app"
	@echo "	make run		--- Run the app"
	@echo "	make build-release	--- Build with optimization flags enabled"
	@echo "	make docker-start	--- Build Dockerfiles and start docker-compose"

build:
	go build -o loadbalancer.out cmd/loadbalancer/main.go

run:
	go run cmd/loadbalancer/main.go

build-release:
	go build -o loadbalancer.out -ldflags "-s -w" cmd/loadbalancer/main.go

docker-start:
	DOCKER_BUILDKIT=1 docker build -t tor-loadbalancer .
	DOCKER_BUILDKIT=1 docker build -t tor-circuit tor-sources/.
	docker-compose up

default: usage
