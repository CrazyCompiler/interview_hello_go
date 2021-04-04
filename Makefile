
build:
	go build -o bin/main ./main/main.go

run:
	go run main/main.go

test:
	go test ./...

test-cov:
	go test ./... -cover

lint:
	go fmt ./...
