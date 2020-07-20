-include .env

build:
	@echo "  >  Building package..."
	go build -o cmd/${BIN_FILENAME} ${GO_PACKAGE_NAME}

proto-gen:
	@echo "  >  Generating gRPC code from protobuf file..."
	protoc -I proto/ proto/route.proto --go_out=plugins=grpc:proto/route

run:
	@echo "  >  Running package..."
	go run ${GO_PACKAGE_NAME}

docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

docker-run:
	docker run -dp 9090:9090 ${DOCKER_IMAGE_NAME}

docker-run-dev:
	docker run -it -p 9090:9090 ${DOCKER_IMAGE_NAME}
