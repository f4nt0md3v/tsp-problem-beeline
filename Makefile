-include .env

BIN_FILENAME=server
DOCKER_IMAGE_NAME=grpc/f4nt0md3v
GO_PACKAGE_NAME=github.com/f4nt0md3v/${PROJECT_NAME}
PROJECT_NAME=tsp-problem-beeline

init-project:
	go mod init ${GO_PACKAGE_NAME}
	@mkdir cmd internal pkg cmd/${PROJECT_NAME} cmd/${PROJECT_NAME}/v1 cmd/${PROJECT_NAME}/internal
	@> main.go
	@> LICENSE
	@> README.md
	@> .env
	@> .gitignore
	@echo "package main\n\nfunc main() {\n\t\n}" > main.go

build:
	@echo "  >  Building package..."
	go build -o cmd/${BIN_FILENAME} ${GO_PACKAGE_NAME}

proto-gen:
	@echo "  >  Generating gRPC code from protobuf file..."
	protoc -I proto/ proto/route.proto --go_out=plugins=grpc:proto/route

run:
	@echo "  >  Running package..."
	go run ${GO_PACKAGE_NAME}

detect-race:
	@echo "  >  Running package in race detection mode..."
	go run -race ${GO_PACKAGE_NAME}

test:
	@echo "  >  Testing package..."
	go test ${GO_PACKAGE_NAME}

fmt:
	@echo "  >  Formating package..."
	go fmt ${GO_PACKAGE_NAME}

clean:
	@echo "  >  Cleaning up build artifacts..."
	go clean

compile-cross:
	@echo "  >  Building package binaries for different platforms..."
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o cmd/${PROJECT_NAME}-freebsd-386 ${GO_PACKAGE_NAME}
	# MacOS
	GOOS=darwin GOARCH=386 go build -o cmd/${PROJECT_NAME}-darwin-386 ${GO_PACKAGE_NAME}
	# Linux
	GOOS=linux GOARCH=386 go build -o cmd/${PROJECT_NAME}-linux-386 ${GO_PACKAGE_NAME}
	# Windows
	GOOS=windows GOARCH=386 go build -o cmd/${PROJECT_NAME}-windows-386 ${GO_PACKAGE_NAME}
        # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o cmd/${PROJECT_NAME}-freebsd-amd64 ${GO_PACKAGE_NAME}
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o cmd/${PROJECT_NAME}-darwin-amd64 ${GO_PACKAGE_NAME}
	# Linux
	GOOS=linux GOARCH=amd64 go build -o cmd/${PROJECT_NAME}-linux-amd64 ${GO_PACKAGE_NAME}
	# Windows
	GOOS=windows GOARCH=amd64 go build -o cmd/${PROJECT_NAME}-windows-amd64 ${GO_PACKAGE_NAME}

lint:
	golangci-lint run ./...

visualize:
	go-callvis ${GO_PACKAGE_NAME}

docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

docker-run:
	docker run -dp 9090:9090 ${DOCKER_IMAGE_NAME}

docker-run-dev:
	docker run -it -p 9090:9090 ${DOCKER_IMAGE_NAME}
