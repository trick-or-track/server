build:
	@go build -o bin/server

run: build
	@go run .

test:
	@go test -v ./...