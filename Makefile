# Variables
GO_BIN_OUT_AUTH := .build/auth
GO_BIN_OUT_DATABASE := .build/database
GO_BIN_OUT_WATERMARK := .build/watermark
DOCKER_IMAGE_AUTH := myapp/auth-service
DOCKER_IMAGE_DATABASE := myapp/database-service
DOCKER_IMAGE_WATERMARK := manansainigits/watermark-service
DOCKER_TAG := latest 

# gRPC code generation for watermark service
codegen-watermark:
	protoc -I="api/pb/watermark" "api/pb/watermark/watermarksvc.proto" --go_out=paths=source_relative:"api/pb/watermark" --go_opt=Mwatermarksvc.proto=/api/pb/watermark
	protoc -I="api/pb/watermark" "api/pb/watermark/watermarksvc.proto" --go-grpc_out=paths=source_relative:"api/pb/watermark" --go-grpc_opt=Mwatermarksvc.proto=/api/pb/watermark

# Common commands
fmt:
	go fmt ./...

clean:
	rm -rf .build

# Build commands for each service
build-auth:
	go build -o $(GO_BIN_OUT_AUTH) ./cmd/auth

build-database:
	go build -o $(GO_BIN_OUT_DATABASE) ./cmd/database

build-watermark: 
	go build -o $(GO_BIN_OUT_WATERMARK) ./cmd/watermark

# Test commands
test:
	go test ./... -v

# Docker build and push commands for each service
dockerbuild-auth:
	docker build -t $(DOCKER_IMAGE_AUTH):$(DOCKER_TAG) -f cmd/auth/Dockerfile .

dockerbuild-database:
	docker build -t $(DOCKER_IMAGE_DATABASE):$(DOCKER_TAG) -f cmd/database/Dockerfile .

dockerbuild-watermark:
	docker build -t $(DOCKER_IMAGE_WATERMARK):$(DOCKER_TAG) -f cmd/watermark/Dockerfile .

dockerpush-auth: dockerbuild-auth
	docker push $(DOCKER_IMAGE_AUTH):$(DOCKER_TAG)

dockerpush-database: dockerbuild-database
	docker push $(DOCKER_IMAGE_DATABASE):$(DOCKER_TAG)

dockerpush-watermark: dockerbuild-watermark
	docker push $(DOCKER_IMAGE_WATERMARK):$(DOCKER_TAG)

# Run commands
run-auth:
	go run ./cmd/auth

run-database:
	go run ./cmd/database

run-watermark:
	go run ./cmd/watermark
