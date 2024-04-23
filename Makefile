TIMESTAMP=$(shell date +%Y%m%d%H%M%S)

build:
	docker buildx build --platform linux/amd64 --output type=docker -t danielian1121/thsr:$(TIMESTAMP) -f ./deployment/thsr/Dockerfile .