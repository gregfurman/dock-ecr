BINARY_NAME=dock-ecr

build:
	GOARCH=amd64 GOOS=darwin go build -o ./dist/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./dist/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./dist/${BINARY_NAME}-windows main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm -rf ./dist

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