BINARY_NAME=dock-ecr

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -mod=mod -o ./dist/${BINARY_NAME}-darwin main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -mod=mod -o ./dist/${BINARY_NAME}-linux main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -mod=mod -o ./dist/${BINARY_NAME}-windows main.go

mod:
	go mod download

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm -rf ./dist

mocks:
	mockgen --source=pkg/docker/client.go --destination=pkg/docker/mock_docker/client.go && \
	mockgen --source=pkg/api/service.go --destination=pkg/api/mock_api/service.go && \
	mockgen --source=pkg/aws/ecr/service.go --destination=pkg/aws/ecr/mock_ecr/service.go

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go get .

vet:
	go vet

lint:
	golangci-lint run -v