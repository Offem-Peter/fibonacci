## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


## run/api: run the cmd/api application - which is using fasthttp with redis
.PHONY: run/api
run/api:
	go run ./cmd/api


## run/fasthttp: run the fasthttp application that is not using redis
.PHONY: run/fasthttp
run/fasthttp:
	go run ./cmd/examples/fasthttp

## run/fiber: run the fiber application that is not using redis
.PHONY: run/fiber
run/fiber:
	go run ./cmd/examples/fiber


## run/httprouter: run the httprouter application that is not using redis
.PHONY: run/httprouter
run/httprouter:
	go run ./cmd/examples/httprouter


## run/redis-fasthttp: run the fasthttp application that is using redis
.PHONY: run/redis-fasthttp
run/redis-fasthttp:
	go run ./cmd/examples/redis-fasthttp

## run/redis-fiber: run the fiber application that is using redis
.PHONY: run/redis-fiber
run/redis-fiber:
	go run ./cmd/examples/redis-fiber

## run/redis-httprouter: run the httprouter application that is using redis
.PHONY: run/redis-httprouter
run/redis-httprouter:
	go run ./cmd/examples/redis-httprouter


# Graph
## graph-with-redis: generates graph data for frameworks using redis
.PHONY: graph-with-redis
graph-with-redis:
	python3 graph_with_redis.py 

# Graph
## graph-without-redis: generates graph data for frameworks not using redis
.PHONY: graph-without-redis
graph-without-redis:
	python3 graph_without_redis.py